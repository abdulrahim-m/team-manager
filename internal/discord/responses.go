package discord

import (
	"fmt"

	"github.com/abdulrahim-m/team-manager/internal/github"
	"github.com/bwmarrin/discordgo"
)

func (h *DiscordHandler) sendNewPRAlert(p github.ReviewPayload) {
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

func (h *DiscordHandler) sendDiscordAlert(p github.ReviewPayload) {
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
