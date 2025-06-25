package telegramSender

import (
	"MotivatorBot/entities"
	"MotivatorBot/interfaces"
	"context"
	"fmt"
)

type TelegramSender struct {
	api interfaces.TelegramBotAPI
}

func NewTelegramSender(api interfaces.TelegramBotAPI) *TelegramSender {
	return &TelegramSender{api: api}
}

func (t *TelegramSender) SendMessage(ctx context.Context, quote *entities.Quote) error {
	message := fmt.Sprintf("📖 %s\n\n _%s_ ✍️", quote.Text, quote.Author)

	err := t.api.SendMessage(ctx, message)
	if err != nil {
		return fmt.Errorf("telegram sendMessage error: %w", err)
	}

	return nil
}
