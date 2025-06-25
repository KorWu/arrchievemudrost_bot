package interfaces

import (
	"MotivatorBot/entities"
	"context"
)

type QuotesAPI interface {
	GetRandomQuote(ctx context.Context) (*entities.Quote, error)
}

type TranslationAPI interface {
	Translate(ctx context.Context, quote *entities.Quote, targetLang string) (*entities.Quote, error)
	TranslateRus(ctx context.Context, quote *entities.Quote) (*entities.Quote, error)
}

type TelegramBotAPI interface {
	SendMessage(ctx context.Context, text string) error
}
