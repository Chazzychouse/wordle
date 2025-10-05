package game

import (
	"errors"
	"time"
)

type Game struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Attempts   []Attempt `json:"attempts" gorm:"serializer:json"`
	Username   string    `json:"username"`
	IsGameOver bool      `json:"isGameOver"`
	Solution   string    `json:"solution"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type GameRequest struct {
	Username string `json:"username"`
}

type Attempt struct {
	Number int    `json:"number"`
	Value  string `json:"value"`
}

type AttemptRequest struct {
	GameID   string `json:"gameId"`
	Username string `json:"username"`
	Value    string `json:"value"`
}

// TODO: Return list of previous attempts
type AttemptResponse struct {
	Result     []Correctness `json:"result"`
	IsGameOver bool          `json:"isGameOver"`
}

const MAX_ATTEMPTS int = 6
const LEN_VALID_ATTEMPT int = 5

type Correctness int

const (
	Incorrect Correctness = iota
	Correct
	OutOfPlace
)

var CorrectAttempt = []Correctness{Correct, Correct, Correct, Correct, Correct}

var ErrInvalidGameID = errors.New("INVALID GAME ID")
var ErrTooManyAttempts = errors.New("TOO MANY ATTEMPTS")
var ErrInvalidAttemptValue = errors.New("ATTEMPT MUST BE EXACTLY 5 LETTERS")
var ErrInvalidUser = errors.New("INVALID USER")
var ErrDuplicateAttempts = errors.New("DUPLICATE ATTEMPT")

func ValidateAttempt(req AttemptRequest, game *Game) (AttemptResponse, error) {
	if req.GameID != game.ID {
		return AttemptResponse{}, ErrInvalidGameID
	}
	if req.Username != game.Username {
		return AttemptResponse{}, ErrInvalidUser
	}
	if len(game.Attempts) >= MAX_ATTEMPTS || game.IsGameOver {
		return AttemptResponse{}, ErrTooManyAttempts
	}
	if len(req.Value) != LEN_VALID_ATTEMPT {
		return AttemptResponse{}, ErrInvalidAttemptValue
	}
	for _, att := range game.Attempts {
		if att.Value == req.Value {
			return AttemptResponse{}, ErrDuplicateAttempts
		}
	}
	result, err := CalculateCorrectness(req.Value, game.Solution)
	if err != nil {
		return AttemptResponse{}, err
	}
	return AttemptResponse{
		Result:     result,
		IsGameOver: isGameOver(result, game.Attempts),
	}, nil
}

func forEachPosition(fn func(i int)) {
	for i := 0; i < LEN_VALID_ATTEMPT; i++ {
		fn(i)
	}
}

func CalculateCorrectness(guess string, sol string) ([]Correctness, error) {
	if guess == sol {
		return CorrectAttempt, nil
	}

	result := make([]Correctness, LEN_VALID_ATTEMPT)
	solCount := make(map[byte]int)

	forEachPosition(func(i int) {
		if guess[i] != sol[i] {
			solCount[sol[i]]++
		}
	})

	forEachPosition(func(i int) {
		if guess[i] == sol[i] {
			result[i] = Correct
		}
	})

	forEachPosition(func(i int) {
		if guess[i] == sol[i] {
			return
		}

		if solCount[guess[i]] > 0 {
			result[i] = OutOfPlace
			solCount[guess[i]]--
		} else {
			result[i] = Incorrect
		}
	})

	return result, nil
}

func isGameOver(result []Correctness, attempts []Attempt) bool {
	return len(attempts)+1 >= MAX_ATTEMPTS || isCorrectGuess(result)
}

func isCorrectGuess(result []Correctness) bool {
	if len(result) != LEN_VALID_ATTEMPT {
		return false
	}
	for _, correctness := range result {
		if correctness != Correct {
			return false
		}
	}
	return true
}
