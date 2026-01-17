package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abdulrahim-m/team-manager/internal/bot"
	"github.com/abdulrahim-m/team-manager/internal/config"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config.LoadConfig()

	dg, err := discordgo.New("Bot " + config.GetBotToken())
	if err != nil {
		log.Fatal(err)
	}
	dg.Open()
	defer dg.Close()

	myBot := &bot.BotHandler{
		DiscordSession: dg,
		UserMap:        bot.LoadMembers(),
		ChannelID:      config.GetChannelID(),
		GithubSecret:   config.GetGithubSecret(),
		ManagerID:      config.GetManagerID(),
	}
	http.HandleFunc("/webhook", myBot.HandleWebhook)
	printLogo()
	log.Println("Server started on :53053...")
	err = http.ListenAndServe(":53053", nil)
	log.Fatal(err)
}

func printLogo() {
	fmt.Println(`
    ____________________     _________     _______ 
   /\                   \   /\        \   /       \
  /  \_______     _______\ /  \        \ /\        \
  \  /      /\    \      / \   \        \_/         \  
   \/______/  \    \____/   \   \        _______     \    
           \   \    \        \   \      /      /\     \
            \   \    \        \   \     \_____/  \     \
             \   \    \        \   \     \    \   \     \
              \   \____\        \   \____/     \   \_____\
               \  /    /         \  /   /       \  /     /
                \/____/           \/___/         \/_____/`)
}
