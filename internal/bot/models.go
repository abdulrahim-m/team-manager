package bot

import (
	"encoding/json"
	"log"
	"os"
)

type member struct {
	GithubHandle     string `json:"github_handle"`
	DiscordSnowflake string `json:"discord_snowflake"`
}

func LoadMembers() map[string]string {
	log.Println("Loading members from database...")
	file, err := os.ReadFile("data")
	if err != nil {
		log.Fatal("Failed to load data:", err)
	}

	var members []member
	json.Unmarshal(file, &members)

	memberMap := map[string]string{}
	for _, m := range members {
		memberMap[m.GithubHandle] = m.DiscordSnowflake
	}

	return memberMap
}
