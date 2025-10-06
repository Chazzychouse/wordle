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

	newID := uuid.New().String()

	newGame := &game.Game{
		ID:         newID,
		Username:   userEmail,
		IsGameOver: false,
		Attempts:   []game.Attempt{},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	newSolution := &game.Solution{
		GameID: newID,
		Value:  word.Word,
	}

	if err := gs.db.CreateGame(newGame); err != nil {
		return nil, errors.New("failed to create game")
	}
	if err := gs.db.CreateSolution(newSolution); err != nil {
		return nil, errors.New("failed to create solution")
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

	return game, nil
}

func (gs *GameService) MakeAttempt(gameID, userEmail, guess string) (*game.AttemptResponse, error) {
	existingGame, err := gs.GetGame(gameID, userEmail)
	if err != nil {
		return nil, err
	}

	existingSolution, err := gs.db.GetSolution(gameID)
	if err != nil {
		return nil, errors.New("solution not found")
	}

	attemptRequest := game.AttemptRequest{
		GameID:   gameID,
		Username: userEmail,
		Value:    guess,
	}

	response, err := game.ValidateAttempt(attemptRequest, existingGame, existingSolution)
	if err != nil {
		return nil, err
	}

	existingGame.Attempts = append(existingGame.Attempts, game.Attempt{
		Number: len(existingGame.Attempts) + 1,
		Value:  guess,
	})
	existingGame.IsGameOver = response.IsGameOver
	existingGame.UpdatedAt = time.Now()

	if err := gs.db.UpdateGame(existingGame); err != nil {
		return nil, errors.New("failed to update game")
	}

	return &response, nil
}

func (gs *GameService) GetUserGames(userEmail string) ([]game.Game, error) {
	return gs.db.GetGamesByUsername(userEmail)
}
