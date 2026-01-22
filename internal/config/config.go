package config

import (
	"log"
	"os"
)

var CONFIG config

type config struct {
	channelID     string
	discToken     string
	githubSecret  string
	managerDiscID string
	managerTeleID string
	teleToken     string
}

func LoadConfig() {
	log.Println("Loading configurations...")
	CONFIG.channelID = os.Getenv("DISCORD_CHANNEL_ID")
	CONFIG.discToken = os.Getenv("DISCORD_BOT_TOKEN")
	CONFIG.githubSecret = os.Getenv("GITHUB_WEBHOOK_SECRET")
	CONFIG.managerDiscID = os.Getenv("MANAGER_DISCORD_ID")
	CONFIG.managerTeleID = os.Getenv("MANAGER_DISCORD_ID")
	CONFIG.teleToken = os.Getenv("TELEGRAM_BOT_TOKEN")
}

func GetChannelID() string     { return CONFIG.channelID }
func GetDiscBotToken() string  { return CONFIG.discToken }
func GetGithubSecret() string  { return CONFIG.githubSecret }
func GetManagerDiscID() string { return CONFIG.managerDiscID }
func GetManagerTeleID() string { return CONFIG.managerDiscID }
func GetTeleBotToken() string  { return CONFIG.teleToken }
