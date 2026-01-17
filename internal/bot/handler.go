package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/abdulrahim-m/team-manager/internal/github"
	"github.com/bwmarrin/discordgo"
)

type BotHandler struct {
	DiscordSession *discordgo.Session
	UserMap        map[string]string
	ChannelID      string
	GithubSecret   string
	ManagerID      string
}

func (h *BotHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	if !github.VerifySignature(body, r.Header.Get("X-Hub-Signature-256")) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var payload github.ReviewPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if payload.Action == "submitted" && payload.Review.State == "changes_requested" {
		h.sendDiscordAlert(payload)
	}

	if payload.Action == "opened" && payload.PullRequest.Base.Ref == "main" {
		h.sendNewPRAlert(payload)
		return
	}
}

func (h *BotHandler) sendDiscordAlert(p github.ReviewPayload) {
	discordID, exists := h.UserMap[p.PullRequest.User.Login]

	var mention = p.PullRequest.User.Login
	if exists {
		mention = fmt.Sprintf("<@%s>", discordID)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🚨 Changes Requested on PR #" + fmt.Sprintf("%d", p.PullRequest.Number),
		Description: fmt.Sprintf("%s, your attention is needed!", mention),
		URL:         p.PullRequest.HTMLURL,
		Color:       0xFF000,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Repo",
				Value:  p.Repository.FullName,
				Inline: true,
			},
			{
				Name:   "Reviewer",
				Value:  p.Sender.Login,
				Inline: true,
			},
			{
				Name:   "Feedback",
				Value:  p.Review.Body,
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Engineering Association Bot • Phase 2",
		},
	}

	h.DiscordSession.ChannelMessageSendEmbed(h.ChannelID, embed)
}

func (h *BotHandler) sendNewPRAlert(p github.ReviewPayload) {
	embed := &discordgo.MessageEmbed{
		Title:       "📥 New Pull Request to Main",
		Description: fmt.Sprintf("Review needed! <@%s>", h.ManagerID),
		URL:         p.PullRequest.HTMLURL,
		Color:       0x0099FF,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Author",
				Value:  p.PullRequest.User.Login,
				Inline: true,
			},
			{
				Name:   "Branch",
				Value:  fmt.Sprintf("%s ➝ main", p.PullRequest.Head.Ref),
				Inline: true,
			},
			{
				Name:   "Title",
				Value:  p.PullRequest.Title,
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Code Review Required",
		},
	}

	h.DiscordSession.ChannelMessageSendEmbed(h.ChannelID, embed)
}
