package internal

import (
	"fmt"
	auth "github.com/Floor-Gang/authclient"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	config  Config
	client  *discordgo.Session
	confLoc string
	auth    auth.AuthClient
}

func Start(config Config, configLocation string) {
	// Setup Discord
	client, _ := discordgo.New("Bot " + config.Token)

	// This is required
	intents := discordgo.MakeIntent(discordgo.IntentsGuildMembers + discordgo.IntentsGuildMessages)
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
	fmt.Println(fmt.Sprintf("Reaction added %s", reaction.MessageID))
}

func (b *Bot) onReactionRemove(_ *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	fmt.Println(fmt.Sprintf("Reaction removed %s", reaction.MessageID))
}

func (b Bot) reply(event *discordgo.MessageCreate, context string) (*discordgo.Message, error) {
	return b.client.ChannelMessageSend(
		event.ChannelID,
		fmt.Sprintf("<@%s> %s", event.Author.ID, context),
	)
}

func (b Bot) addReaction(message *discordgo.Message, emoij string) error {
	return b.client.MessageReactionAdd(message.ChannelID, message.ID, ":emoij:" + emoij)
}