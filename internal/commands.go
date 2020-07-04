package internal

import (
	"fmt"
	util "github.com/Floor-Gang/utilpkg"
	dg "github.com/bwmarrin/discordgo"
	"strings"
)

func (b *Bot) onMessage(_ *dg.Session, message *dg.MessageCreate) {
	if len(message.GuildID) == 0 || !strings.HasPrefix(message.Content, b.config.Prefix) {
		return
	}

	args := strings.Split(message.Content, " ")

	if len(args) < 3 {
		return
	}

	// auth, err := b.auth.Auth(message.Author.ID)

	//if err != nil {
	//	log.Printf("Failed to authentication \"%s\", because\n%s", message.Author.ID, err.Error())
	//}

	switch args[1] {
	case "add":
		// Make sure only actual relevant arguments get passed along
		args = removeFromSlice("add", append(args[1:]))

		if len(args) >= 3 && len(args) <= 4 {
			messageResponse := b.addRole(args, message.Author.ID, message.GuildID, message.ChannelID)
			_, _ = b.reply(message, messageResponse)
		} else {
			_, _ = b.reply(message, fmt.Sprintf("**Invalid syntax** ``%s add (channel*) [message id] [emoij] [role]``", b.config.Prefix))
		}

		break
	case "remove":
		// Make sure only actual relevant arguments get passed along
		args = removeFromSlice("add", append(args[1:]))

		if len(args) >= 2 && len(args) <= 3 {
			_ = b.removeRole(args, message.GuildID, message.ChannelID)
		} else {
			_, _ = b.reply(message, fmt.Sprintf("**Invalid syntax** ``%s remove (channel*) [message id] [emoij]``", b.config.Prefix))
		}

		break
	}
}

func (b *Bot) addRole(args []string, userID string, guildID string, channelID string) string {
	defaultErrorMessage := fmt.Sprintf("**Invalid syntax** ``%s add (channel*) [message id] [emoij] [role]``", b.config.Prefix)

	if len(args) == 3 {
		Message, err := b.client.ChannelMessage(channelID, args[0])

		if err != nil {
			// Invalid messageID
			fmt.Println(err)
			return defaultErrorMessage
		}

		_ = b.addReaction(Message, util.FilterTag(args[1]))
		err = b.db.createRoleReaction(guildID, channelID, Message.ID, util.FilterTag(args[1]), util.FilterTag(args[2]))

		if err != nil {
			// Couldn't save data to SQLite.
			return "An error occurred while saving the data, contact a developer."
		}
	} else {
		Channel, err := b.client.Channel(util.FilterTag(args[0]))

		if err != nil {
			// Channel is invalid
			fmt.Println(err)
			return defaultErrorMessage
		}

		Message, err := b.client.ChannelMessage(Channel.ID, args[1])

		if err != nil {
			// Message wasn't found, probably an invalid combination of channel & message
			return "Message wasn't found in mentioned channel."
		}

		_ = b.addReaction(Message, util.FilterTag(args[2]))
		err = b.db.createRoleReaction(guildID, Channel.ID, Message.ID, util.FilterTag(args[2]), util.FilterTag(args[3]))

		if err != nil {
			// Couldn't save data to SQLite.
			return "An error occurred while saving the data, contact a developer."
		}
	}

	return "added role-reaction to successfully!"
}

func (b *Bot) removeRole(args []string, guildID string, channelID string) string {
	defaultErrorMessage := fmt.Sprintf("**Invalid syntax** ``%s remove (channel*) [message id] [emoij]``", b.config.Prefix)

	if len(args) == 2 {
		Message, err := b.client.ChannelMessage(channelID, args[0])

		if err != nil {
			fmt.Println(err)
			return defaultErrorMessage
		}

		err = b.db.removeRoleReaction(guildID, channelID, Message.ID, util.FilterTag(args[1]))

		if err != nil {
			fmt.Println(err)
		}

		err = b.removeReaction(Message, util.FilterTag(args[1]))

		if err != nil {
			fmt.Println(err)
		}
	} else {
		Channel, err := b.client.Channel(util.FilterTag(args[0]))

		if err != nil {
			// Channel is invalid
			fmt.Println(err)
			return defaultErrorMessage
		}

		Message, err := b.client.ChannelMessage(Channel.ID, args[1])

		if err != nil {
			fmt.Println(err)
			return defaultErrorMessage
		}

		err = b.db.removeRoleReaction(Channel.GuildID, Channel.ID, Message.ID, util.FilterTag(args[2]))

		if err != nil {
			fmt.Println(err)
		}

		err = b.removeReaction(Message, util.FilterTag(args[2]))

		if err != nil {
			fmt.Println(err)
		}
	}

	return "Removed the emoij from the message"
}