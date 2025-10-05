package services

import (
	"errors"
	"time"
	"wordle/internal/database"
	"wordle/internal/game"

	"github.com/google/uuid"
)

type GameService struct {
	db          *database.Database
	wordService *WordService
}

func NewGameService(db *database.Database, wordService *WordService) *GameService {
	return &GameService{
		db:          db,
		wordService: wordService,
	}
}

func (gs *GameService) CreateGame(userEmail string) (*game.Game, error) {
	word, err := gs.wordService.GetRandomWord()
	if err != nil {
		return nil, errors.New("failed to get random word")
	}

	newGame := &game.Game{
		ID:         uuid.New().String(),
		Username:   userEmail,
		Solution:   word.Word,
		IsGameOver: false,
		Attempts:   []game.Attempt{},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := gs.db.CreateGame(newGame); err != nil {
		return nil, errors.New("failed to create game")
	}

	return newGame, nil
}

func (gs *GameService) GetGame(gameID, userEmail string) (*game.Game, error) {
	game, err := gs.db.GetGame(gameID)
	if err != nil {
		return nil, errors.New("game not found")
	}

	if game.Username != userEmail {
		return nil, errors.New("access denied")
	}

	if !game.IsGameOver {
		game.Solution = "*****"
	}

	return game, nil
}

func (gs *GameService) MakeAttempt(gameID, userEmail, guess string) (*game.AttemptResponse, error) {
	existingGame, err := gs.GetGame(gameID, userEmail)
	if err != nil {
		return nil, err
	}

	attemptRequest := game.AttemptRequest{
		GameID:   gameID,
		Username: userEmail,
		Value:    guess,
	}

	response, err := game.ValidateAttempt(attemptRequest, existingGame)
	if err != nil {
		return nil, err
	}

	existingGame.Attempts = append(existingGame.Attempts, game.Attempt{
		Number: len(existingGame.Attempts) + 1,
		Value:  guess,
	})
	existingGame.IsGameOver = response.IsGameOver
	existingGame.UpdatedAt = time.Now()

	// Save updated game
	if err := gs.db.UpdateGame(existingGame); err != nil {
		return nil, errors.New("failed to update game")
	}

	return &response, nil
}

// GetUserGames retrieves all games for a user
func (gs *GameService) GetUserGames(userEmail string) ([]game.Game, error) {
	return gs.db.GetGamesByUsername(userEmail)
}
