package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	invalidUrlError = errors.New("url is invalid")
)

// This func to check erros and send user-friendly message to bot chat
func (b *Bot) handleError(chatId int64, err error) {
	var errorMessage string

	switch err {
	case invalidUrlError:
		errorMessage = b.errors.InvalidURL
	default:
		errorMessage = b.errors.Default
	}

	msg := tgbotapi.NewMessage(chatId, errorMessage)
	b.bot.Send(msg)
}
