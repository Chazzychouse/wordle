package handlers

import (
	"encoding/json"
	"net/http"
	"wordle/internal/database"
	"wordle/internal/services"

	"github.com/gorilla/mux"
)

type GameHandler struct {
	gameService *services.GameService
}

func NewGameHandler() *GameHandler {
	db := database.NewDatabase()
	wordService := services.NewWordService()
	gameService := services.NewGameService(db, wordService)

	return &GameHandler{
		gameService: gameService,
	}
}

func (gh *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	// userEmail := middleware.GetUserEmail(r.Context())
	// if userEmail == "" {
	// 	http.Error(w, "User not authenticated", http.StatusUnauthorized)
	// 	return
	// }

	game, err := gh.gameService.CreateGame("jerry@example.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"game_id": game.ID,
		"message": "Game created successfully",
	})
}

func (gh *GameHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	// TODO: Protect this when OAUTH is implemented
	// userEmail := middleware.GetUserEmail(r.Context())
	// if userEmail == "" {
	// 	http.Error(w, "User not authenticated", http.StatusUnauthorized)
	// 	return
	// }

	gameId := mux.Vars(r)["id"]
	if gameId == "" {
		http.Error(w, "Game ID is required", http.StatusBadRequest)
		return
	}

	game, err := gh.gameService.GetGame(gameId, "jerry@example.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"game": game,
	})
}

func (gh *GameHandler) MakeAttempt(w http.ResponseWriter, r *http.Request) {
	// TODO: Protect this when OAUTH is implemented
	// userEmail := middleware.GetUserEmail(r.Context())
	// if userEmail == "" {
	// 	http.Error(w, "User not authenticated", http.StatusUnauthorized)
	// 	return
	// }

	gameId := mux.Vars(r)["id"]
	if gameId == "" {
		http.Error(w, "Game ID is required", http.StatusBadRequest)
		return
	}

	var request struct {
		Guess string `json:"guess"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := gh.gameService.MakeAttempt(gameId, "jerry@example.com", request.Guess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result":     response.Result,
		"isGameOver": response.IsGameOver,
	})
}

func (gh *GameHandler) GetUserGames(w http.ResponseWriter, r *http.Request) {
	// TODO: Protect this when OAUTH is implemented
	// userEmail := middleware.GetUserEmail(r.Context())
	// if userEmail == "" {
	// 	http.Error(w, "User not authenticated", http.StatusUnauthorized)
	// 	return
	// }

	userEmail := mux.Vars(r)["email"]
	if userEmail == "" {
		http.Error(w, "User email is required", http.StatusBadRequest)
		return
	}
	games, err := gh.gameService.GetUserGames(userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"games": games,
	})
}
