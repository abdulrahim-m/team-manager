package discord

import (
	"log"

	"github.com/abdulrahim-m/team-manager/internal/config"
	"github.com/bwmarrin/discordgo"
)

func SendNotification(message string) {
	dg, err := discordgo.New("Bot " + config.GetBotToken())
	if err != nil {
		log.Println("Error creating Discord session:", err)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Println("Error creating Discord session:", err)
		return
	}
	defer dg.Close()

	dg.ChannelMessageSend(config.GetChannelID(), message)
}
