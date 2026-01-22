package discord

import "github.com/bwmarrin/discordgo"

func (h *DiscordHandler) sendStartMessage() {
	embed := &discordgo.MessageEmbed{
		Title: "Github notification relay enabled",
		// Description: "",
		Color: 0x00ff00,
	}
	h.DiscordSession.ChannelMessageSendEmbed(h.ChannelID, embed)
}
