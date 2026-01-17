package config

import (
	"log"
	"os"
)

var CONFIG config

type config struct {
	channelID    string
	token        string
	githubSecret string
	managerID    string
}

func LoadConfig() {
	log.Println("Loading configurations...")
	CONFIG.channelID = os.Getenv("DISCORD_CHANNEL_ID")
	CONFIG.token = os.Getenv("DISCORD_BOT_TOKEN")
	CONFIG.githubSecret = os.Getenv("GITHUB_WEBHOOK_SECRET")
	CONFIG.managerID = os.Getenv("MANAGER_DISCORD_ID")
}

func GetChannelID() string    { return CONFIG.channelID }
func GetBotToken() string     { return CONFIG.token }
func GetGithubSecret() string { return CONFIG.githubSecret }
func GetManagerID() string    { return CONFIG.managerID }
