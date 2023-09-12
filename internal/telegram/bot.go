package telegram

import (
	"salesforceanton/news-tg-bot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI

	errors    config.Errors
	responses config.Responses
}

func NewBot(cfg *config.Config) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.BotToken)
	bot.Debug = true

	if err != nil {
		return nil, err
	}
	return &Bot{
		bot:       bot,
		errors:    cfg.Messages.Errors,
		responses: cfg.Messages.Responses,
	}, nil
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Run cycle in goroutine to ping tg API for updates
	updates := b.bot.GetUpdatesChan(u)

	// Handle Updates
	for update := range updates {
		// If we got a message
		if update.Message == nil {
			continue
		}

		chatId := update.Message.Chat.ID
		// Handle Command
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(chatId, err)
			}
			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(chatId, err)
		}
	}
}
