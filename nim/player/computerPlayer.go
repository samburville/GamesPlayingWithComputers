package player

import (
	"bufio"
	"fmt"

	"github.com/samburville/GamesPlayingWithComputers/nim/board"
)

// ComputerPlayer structure for a computer player of nim
type ComputerPlayer struct {
}

// PlayMove plays a computer players move and returns true if successful
func (c ComputerPlayer) PlayMove(gameState *board.GameState, reader *bufio.Reader) bool {

	if gameState.IsGameover() {
		panic("Does Not Compute")
	}

	pile, matchesToTake := calculateOptimalMove(gameState)

	fmt.Printf("Computer removes from %d matches from pile %d\n", matchesToTake, pile)

	gameState.PlayMove(pile, matchesToTake)

	return true
}

func calculateOptimalMove(gameState *board.GameState) (int, int) {

	gameBoard := gameState.GetGameBoard()

	nimSum := calculateNimSum(gameBoard)

	if nimSum == 0 {

		for i := range gameBoard {
			if gameBoard[i] > 0 {

				return i + 1, 1
			}
		}
	}

	for i := range gameBoard {
		targetSize := gameBoard[i] ^ nimSum

		if (targetSize) < gameBoard[i] {

			return i + 1, gameBoard[i] - (targetSize)
		}
	}

	panic("This shouldn't happen")
}

func calculateNimSum(gameBoard []int) int {
	nimSum := 0

	for i := 0; i < len(gameBoard); i++ {
		nimSum = nimSum ^ gameBoard[i]
	}

	return nimSum
}
