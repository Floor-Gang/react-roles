package internal

import (
	"fmt"
	auth "github.com/Floor-Gang/authclient"
	"github.com/bwmarrin/discordgo"
	"log"
	"regexp"
)

type Bot struct {
	config  Config
	client  *discordgo.Session
	confLoc string
	auth    auth.AuthClient
	db      Controller
}

func Start(config Config, configLocation string) {
	// Setup the database
	database := GetController(config.DBLocation)

	// Setup Discord
	client, _ := discordgo.New("Bot " + config.Token)

	// This is required
	intents := discordgo.MakeIntent(discordgo.IntentsGuildMembers + discordgo.IntentsGuildMessages + discordgo.IntentsGuildMessageReactions)
	client.Identify.Intents = intents

	// Setup Authentication client
	authClient, err := auth.GetClient(config.Auth)

	if err != nil {
		// Ignore auth
		// log.Fatalln("Failed to connect to authentication server because \n" + err.Error())
	}

	bot := Bot{
		config:  config,
		client:  client,
		confLoc: configLocation,
		auth:    authClient,
		db:      database,
	}

	// Add event listeners
	client.AddHandler(bot.onReady)
	client.AddHandler(bot.onReactionAdd)
	client.AddHandler(bot.onReactionRemove)
	client.AddHandler(bot.onMessage)

	if err = client.Open(); err != nil {
		log.Fatalln("Failed to connect to Discord, was an access token provided?\n" + err.Error())
	}
}

func (b *Bot) onReady(_ *discordgo.Session, ready *discordgo.Ready) {
	log.Printf("Ready as %s#%s\n", ready.User.Username, ready.User.Discriminator)
}

func (b *Bot) onReactionAdd(_ *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	data, err := b.db.getAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, data := range data {
		compareEmoji := (regexp.MustCompile(`:(.*?):`)).ReplaceAllString(data.reaction, "")

		if data.messageID != reaction.MessageID && data.channelID != reaction.ChannelID {
			continue
		}

		if data.GuildID == reaction.GuildID && reaction.Emoji.ID == compareEmoji {
			err = b.client.GuildMemberRoleAdd(reaction.GuildID, reaction.UserID, data.role)

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (b *Bot) onReactionRemove(_ *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	data, err := b.db.getAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, data := range data {
		compareEmoji := (regexp.MustCompile(`:(.*?):`)).ReplaceAllString(data.reaction, "")

		if data.messageID == reaction.MessageID && data.channelID == reaction.ChannelID {
			continue
		}

		if data.GuildID == reaction.GuildID && reaction.Emoji.ID == compareEmoji {
			err = b.client.GuildMemberRoleRemove(reaction.GuildID, reaction.UserID, data.role)

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// TODO: Use the util package instead of this
func (b Bot) reply(event *discordgo.MessageCreate, context string) (*discordgo.Message, error) {
	return b.client.ChannelMessageSend(
		event.ChannelID,
		fmt.Sprintf("<@%s> %s", event.Author.ID, context),
	)
}

func (b Bot) addReaction(message *discordgo.Message, emoij string) error {
	return b.client.MessageReactionAdd(message.ChannelID, message.ID, emoij)
}

func (b Bot) removeReaction(message *discordgo.Message, emoij string) error {
	// Getting every user that reacted with this specific emoij, up to 5 rate-limit. And it's not suposed to be used when a 1000 people reacted anyways.
	data, err := b.client.MessageReactions(message.ChannelID, message.ID, emoij, 5, "", "")

	if err != nil {
		return err
	}

	for _, item := range data {
		// Remove their reaction.
		err = b.client.MessageReactionRemove(message.ChannelID, message.ID, emoij, item.ID)

		if err != nil {
			return err
		}
	}

	return err
}
