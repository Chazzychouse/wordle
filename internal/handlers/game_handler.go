package handlers

import (
	"fmt"
	"net/http"
	"wordle/internal/services"
)

type GameHandler struct {
	wordService *services.WordService
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		wordService: services.NewWordService(),
	}
}

func (gh *GameHandler) TestNewWord(w http.ResponseWriter, r *http.Request) {
	word, err := gh.wordService.GetRandomWord()
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		http.Error(w, "Failed to get word", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Random Word: %s\n", word)
}
