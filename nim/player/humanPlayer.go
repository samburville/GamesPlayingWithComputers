package player

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/samburville/GamesPlayingWithComputers/nim/board"
	"github.com/samburville/GamesPlayingWithComputers/nim/consolehelper"
)

// HumanPlayer structure for a human player of nim
type HumanPlayer struct {
}

// PlayMove plays a human players move and returns true if successful
func (h HumanPlayer) PlayMove(gameState *board.GameState, reader *bufio.Reader) bool {

	pile, err := consolehelper.AskPlayerForInt("Which pile would you like to remove from?", reader)

	if err != nil {
		fmt.Println("Invalid User input. Try Again")
		return false
	}

	if pile > len(gameState.GameBoard) {
		fmt.Println("Pile is out of range of the game board. Try Again")
		fmt.Println()
		return false
	}

	matches, err := consolehelper.AskPlayerForInt("How many matches would you like to remove?", reader)

	if err != nil {
		fmt.Println("Invalid User input. Try Again")
		return false
	}

	matchesLeftover := len(gameState.GameBoard[pile-1]) - matches

	if matchesLeftover < 0 {
		fmt.Println("Number of matches to remove is larger than the pile. Try Again")
		fmt.Println()
		return false
	}

	gameState.GameBoard[pile-1] = strings.Repeat("|", matchesLeftover)
	return true
}
