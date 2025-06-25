package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	ChatID int64
	Token  string
}

func LoadConfig() (*Config, error) {

	token := os.Getenv("BOT_TOKEN")
	chatIDstr := os.Getenv("CHAT_ID")

	if token == "" || chatIDstr == "" {
		log.Fatal("BOT_TOKEN and CHAT_ID must be set")
	}

	chatID, err := strconv.ParseInt(chatIDstr, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("can't parse CHAT_ID: %w", err)
	}

	return &Config{
		ChatID: chatID,
		Token:  token,
	}, nil
}
