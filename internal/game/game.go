package game

import (
	"errors"
)

type Game struct {
	ID         string    `json:id`
	Attempts   []Attempt `json:attempts`
	Username   string    `json:username`
	IsGameOver bool      `json:isGameOver`
	Solution   string    `json:solution`
}

type Attempt struct {
	Number uint16 `json:number`
	Value  string `json:value`
}

type AttemptRequest struct {
	GameID   string `json:gameId`
	Username string `json:username`
	Value    string `json:value`
}

type AttemptResponse struct {
	Result     []Correctness `json:result`
	IsGameOver bool          `json:isGameOver`
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

func ValidateGame(req AttemptRequest, game *Game) (bool, error) {
	if req.GameID != game.ID {
		return false, ErrInvalidGameID
	}
	if req.Username != game.Username {
		return false, ErrInvalidUser
	}
	if len(game.Attempts) >= MAX_ATTEMPTS || game.IsGameOver {
		return false, ErrTooManyAttempts
	}
	for _, att := range game.Attempts {
		if att.Value == req.Value {
			return false, ErrDuplicateAttempts
		}
	}
	return true, nil
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
