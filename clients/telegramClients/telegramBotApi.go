package telegramClients

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBotApi struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func NewTelegramBotApi(token string, chatID int64) (*TelegramBotApi, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("can't create bot: %w", err)
	}

	return &TelegramBotApi{bot: bot, chatID: chatID}, nil
}

func (t *TelegramBotApi) SendMessage(ctx context.Context, message string) error {
	msg := tgbotapi.NewMessage(t.chatID, message)

	_, err := t.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}

	return nil
}
