package main

import (
	"fmt"
	"strings"
)

// Game structure for playing a game of nim
type Game struct {
	difficulty     string
	piles, matches int
}

// GameState structure for playing a game of nim
type GameState struct {
	gameBoard []string
	gameover  bool
}

func main() {
	piles := 4
	matches := 5

	fmt.Println(NewGameState(piles, matches))
}

// NewGameState creates a new game state with a number of piles
// each with the specified number of matches
func NewGameState(piles, matches int) *GameState {

	singlePileString := strings.Repeat("|", matches)

	gameBoard := make([]string, piles)

	for i := 0; i < piles; i++ {
		gameBoard[i] = singlePileString
	}

	isGameover := piles == 1 && matches == 1

	return &GameState{gameBoard, isGameover}
}
