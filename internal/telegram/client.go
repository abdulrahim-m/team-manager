package telegram

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/abdulrahim-m/team-manager/internal/config"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartTelegramBot(ch chan<- string) {
	log.Println("Starting Telegram bot...")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var handler func(ctx context.Context, b *bot.Bot, update *models.Update) = func(ctx context.Context, b *bot.Bot, update *models.Update) {
		botHandler(ctx, b, update, ch)
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(config.GetTeleBotToken(), opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
	log.Println("Telegram bot started and ready.")
}
