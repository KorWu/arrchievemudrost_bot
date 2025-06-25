package telegramClients

import (
	"MotivatorBot/entities"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const ZenQuotesApiUrl = "https://zenquotes.io/api/random"

type ZenQuotesAPI struct{}

func NewZenQuotesAPI() *ZenQuotesAPI {
	return &ZenQuotesAPI{}
}

type responseFromAPI struct {
	Text   string `json:"q"`
	Author string `json:"a"`
}

func (z *ZenQuotesAPI) GetRandomQuote(ctx context.Context) (*entities.Quote, error) {
	resp, err := http.Get(ZenQuotesApiUrl)
	if err != nil {
		return nil, fmt.Errorf("could not get random quote: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	quotes := make([]responseFromAPI, 0)
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %w", err)
	}

	if len(quotes) == 0 {
		return nil, fmt.Errorf("empty response from zenquotes")
	}

	return &entities.Quote{
		Text:   quotes[0].Text,
		Author: quotes[0].Author,
	}, nil
}
