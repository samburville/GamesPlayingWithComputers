package board

import (
	"fmt"
	"strings"
)

// GameState structure for playing a game of nim
type GameState struct {
	gameBoard []string
	gameover  bool
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

// PrintBoard prints the current state of the game
func (gameState *GameState) PrintBoard() {
	fmt.Println("Current Game Board")
	fmt.Println()

	for i := range gameState.gameBoard {
		fmt.Println(gameState.gameBoard[i])
	}

	fmt.Println()
}

// IsGameover checks for gameover based on the number of matches remaining
func (gameState *GameState) IsGameover() bool {
	return 1 == gameState.GetNumberOfMatchesRemaining()
}

// GetNumberOfMatchesRemaining returns the number of matches left in the game
func (gameState *GameState) GetNumberOfMatchesRemaining() int {
	matchesRemaining := 0

	for i := range gameState.gameBoard {
		matchesRemaining = matchesRemaining + len(gameState.gameBoard[i])
	}
	return matchesRemaining
}
