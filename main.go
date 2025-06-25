package main

import (
	"MotivatorBot/Scheduler"
	"MotivatorBot/clients/telegramClients"
	"MotivatorBot/config"
	"MotivatorBot/messageSender/telegramSender"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config")
	}

	quoteService := telegramClients.NewZenQuotesAPI()
	translator := telegramClients.NewMyMemoryApi()
	telegramAPI, err := telegramClients.NewTelegramBotApi(cfg.Token, cfg.ChatID)
	if err != nil {
		log.Fatal("Error creating telegram bot ", err)
	}
	messageSender := telegramSender.NewTelegramSender(telegramAPI)

	eventManager := Scheduler.NewEventManager(translator, quoteService, messageSender)

	eventManager.Start()
}
