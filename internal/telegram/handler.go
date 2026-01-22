package telegram

import (
	"context"
	"os/exec"

	"github.com/abdulrahim-m/team-manager/internal/config"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func botHandler(ctx context.Context, b *bot.Bot, update *models.Update, ch chan<- string) {
	if string(update.Message.Contact.UserID) != config.GetManagerTeleID() {
		return
	}

	switch update.Message.Text {
	case "/start":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text: `
Commands available:
- /start
			`,
		})
	case "/notify":
		ch <- "notify"
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Sent to discord",
		})
	case "/reboot":
		cmd := exec.Command("pm2", "restart", "~/bot/ecosystem.config.js")
		err := cmd.Run()
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
	}
}
