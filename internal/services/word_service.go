package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WordService struct {
	client  *http.Client
	baseUrl string
}

type WordApiResponse struct {
	Word     string `json:"word"`
	Category string `json:"category"`
	Length   int    `json:"length"`
	Language string `json:"language"`
}

type WordApiArrayResponse []WordApiResponse

func NewWordService() *WordService {
	return &WordService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: "https://random-words-api.kushcreates.com/api",
	}
}

func (ws *WordService) GetRandomWord() (WordApiResponse, error) {
	url := fmt.Sprintf("%s?language=en&category=wordle&length=5&words=1", ws.baseUrl)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return WordApiResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := ws.client.Do(req)
	if err != nil {
		return WordApiResponse{}, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WordApiResponse{}, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WordApiResponse{}, fmt.Errorf("error reading response: %w", err)
	}

	var apiResponse WordApiArrayResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return WordApiResponse{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return apiResponse[0], nil
}
