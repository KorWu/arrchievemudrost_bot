package Scheduler

import (
	"MotivatorBot/interfaces"
	"MotivatorBot/messageSender/telegramSender"
	"context"
	"github.com/robfig/cron/v3"
	"log"
)

type EventManager struct {
	cron       *cron.Cron
	translator interfaces.TranslationAPI
	fetcher    interfaces.QuotesAPI
	sender     *telegramSender.TelegramSender
}

func NewEventManager(translator interfaces.TranslationAPI, fetcher interfaces.QuotesAPI, sender *telegramSender.TelegramSender) *EventManager {
	return &EventManager{
		cron:       cron.New(),
		translator: translator,
		fetcher:    fetcher,
		sender:     sender,
	}
}

func (em *EventManager) Start() {

	_, err := em.cron.AddFunc("0 4,8,14,18 * * *", func() {

		ctx := context.Context(context.Background())

		quote, err := em.fetcher.GetRandomQuote(ctx)
		if err != nil {
			log.Printf("can't fetch the message: %v", err)
			return
		}

		translatedQuote, err := em.translator.TranslateRus(ctx, quote)
		if err != nil {
			log.Printf("can't translate the quote: %v", err)
			return
		}

		if err := em.sender.SendMessage(ctx, translatedQuote); err != nil {
			log.Printf("can't send the message: %v", err)
			return
		} else {
			log.Printf("the quote sended: %v", quote)
		}
	})
	if err != nil {
		log.Fatal("can't make the cron task", err)
	}

	em.cron.Start()

	log.Println("cron task started")

	defer em.cron.Stop()

	select {}
}
