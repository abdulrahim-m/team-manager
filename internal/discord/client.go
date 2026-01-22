package discord

import (
	"log"
	"net/http"

	"github.com/abdulrahim-m/team-manager/internal/config"
	"github.com/bwmarrin/discordgo"
)

func StartDiscordBot(ch <-chan string) {
	dg, err := discordgo.New("Bot " + config.GetDiscBotToken())
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
		return
	}
	defer dg.Close()

	b := &DiscordHandler{
		DiscordSession: dg,
		UserMap:        LoadMembers(),
		ChannelID:      config.GetChannelID(),
		GithubSecret:   config.GetGithubSecret(),
		ManagerID:      config.GetManagerDiscID(),
	}

	http.HandleFunc("/webhook", b.HandleWebhook)

	b.HangleTelegram(ch)
}
