package internal

import (
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
		response := b.addRole(args[2], message.Author.ID, message.GuildID)
		messageResponse, _ := b.reply(message, response)
		_ = b.addReaction(messageResponse, "726592398913699872")
		break
	case "remove":
		response := b.removeRole(args[2], message.Author.ID)
		_, _ = b.reply(message, response)
		break
	}
}

func (b *Bot) addRole(roleID string, userID string, guildID string) string {
	return "role added!"
}

func (b *Bot) removeRole(roleID string, userID string) string {
	return "role removed!"
}
