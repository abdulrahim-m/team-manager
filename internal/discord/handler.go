package discord

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/abdulrahim-m/team-manager/internal/github"
	"github.com/abdulrahim-m/team-manager/storage"
	"github.com/bwmarrin/discordgo"
)

type DiscordHandler struct {
	DiscordSession *discordgo.Session
	UserMap        map[string]string
	ChannelID      string
	GithubSecret   string
	ManagerID      string
}

func (h *DiscordHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	storage.SaveGithubLog(body)

	// if !github.VerifySignature(body, r.Header.Get("X-Hub-Signature-256")) {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	var payload github.ReviewPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("Bad Request:", err)
		return
	}

	if payload.Action == "submitted" && payload.Review.State == "changes_requested" {
		log.Println("submitted, sending response")
		h.sendDiscordAlert(payload)
		log.Println("response sent.")
		return
	}

	if payload.Action == "opened" && payload.PullRequest.Base.Ref == "main" {
		log.Println("opened, sending response")
		h.sendNewPRAlert(payload)
		log.Println("response sent.")
		return
	}
	log.Println("format not matched, aporting")
}

func (h *DiscordHandler) HangleTelegram(ch <-chan string) {
	log.Println("Discord bot Handling Telegram commands...")
	for command := range ch {
		switch command {
		case "notify":
			h.sendStartMessage()
		}
	}
}
