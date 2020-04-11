package player

import (
	"bufio"
	"fmt"

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

	matches, err := consolehelper.AskPlayerForInt("How many matches would you like to remove?", reader)

	if err != nil {
		fmt.Println("Invalid User input. Try Again")
		return false
	}

	success := gameState.PlayMove(pile, matches)

	if !success {
		fmt.Printf("%d matches from pile %d is an invalid move.\n", matches, pile)
		return false
	}

	return true
}
