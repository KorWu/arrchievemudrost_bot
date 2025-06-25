package telegramClients

import (
	"MotivatorBot/entities"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	MyMemoryUrl = "https://api.mymemory.translated.net/get"
)

type MyMemoryApi struct{}

func NewMyMemoryApi() *MyMemoryApi {
	return &MyMemoryApi{}
}

type responseFromApi struct {
	ResponseData struct {
		TranslatedText string `json:"translatedText"`
	} `json:"responseData"`
}

func (m *MyMemoryApi) Translate(ctx context.Context, quote *entities.Quote, targetLang string) (*entities.Quote, error) {
	u := MyMemoryUrl
	params := url.Values{}

	params.Set("q", quote.Text)
	params.Set("langpair", "en|"+targetLang)

	resp, err := http.Get(u + "?" + params.Encode())
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error making http request to MyMemoryApi: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	defer resp.Body.Close()

	translatedText := responseFromApi{}
	if err := json.Unmarshal(body, &translatedText); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	if translatedText.ResponseData.TranslatedText == "" {
		return nil, fmt.Errorf("translated text is empty")
	}

	params.Set("q", quote.Author)

	resp, err = http.Get(u + "?" + params.Encode())
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error making http request to MyMemoryApi: %w", err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	defer resp.Body.Close()

	translatedAuthor := responseFromApi{}
	if err := json.Unmarshal(body, &translatedAuthor); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	if translatedAuthor.ResponseData.TranslatedText == "" {
		return nil, fmt.Errorf("translated author is empty")
	}

	return &entities.Quote{
		Text:   translatedText.ResponseData.TranslatedText,
		Author: translatedAuthor.ResponseData.TranslatedText,
	}, nil
}

func (m *MyMemoryApi) TranslateRus(ctx context.Context, quote *entities.Quote) (*entities.Quote, error) {
	return m.Translate(ctx, quote, "ru")
}
