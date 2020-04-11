package board

import (
	"fmt"
	"strings"
)

// GameState structure for playing a game of nim
type GameState struct {
	gameBoard []int
}

// NewGameState creates a new game state with a number of piles
// each with the specified number of matches
func NewGameState(piles, matches int) *GameState {

	gameBoard := make([]int, piles)

	for i := 0; i < piles; i++ {
		gameBoard[i] = matches
	}

	return &GameState{gameBoard}
}

// PlayMove removes X matches from pile Y. If this is not a valid move, then it will return false
func (gameState *GameState) PlayMove(pileToTakeFrom, matchesToTake int) bool {

	if pileToTakeFrom < 1 || pileToTakeFrom > len(gameState.gameBoard) {
		return false
	}

	if matchesToTake <= 0 || matchesToTake > gameState.gameBoard[pileToTakeFrom-1] {
		return false
	}

	matchesLeftover := gameState.gameBoard[pileToTakeFrom-1] - matchesToTake

	gameState.gameBoard[pileToTakeFrom-1] = matchesLeftover

	return true
}

// PrintBoard prints the current state of the game
func (gameState *GameState) PrintBoard() {

	for i := range gameState.gameBoard {
		fmt.Println(strings.Repeat("|", gameState.gameBoard[i]))
	}

	fmt.Println()
}

// IsGameover checks for gameover based on the number of matches remaining
func (gameState *GameState) IsGameover() bool {
	return 0 == gameState.GetNumberOfMatchesRemaining()
}

// GetNumberOfMatchesRemaining returns the number of matches left in the
func (gameState *GameState) GetNumberOfMatchesRemaining() int {
	matchesRemaining := 0

	for i := range gameState.gameBoard {
		matchesRemaining = matchesRemaining + gameState.gameBoard[i]
	}

	return matchesRemaining
}

// GetGameBoard returns a new object which represents the current game board state
func (gameState *GameState) GetGameBoard() []int {

	return gameState.gameBoard
}
