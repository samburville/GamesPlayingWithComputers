package player

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/samburville/GamesPlayingWithComputers/nim/board"
)

// ComputerPlayer structure for a computer player of nim
type ComputerPlayer struct {
}

// PlayMove plays a computer players move and returns true if successful
func (c ComputerPlayer) PlayMove(gameState *board.GameState, reader *bufio.Reader) bool {

	if gameState.Gameover {
		panic("Does Not Compute")
	}

	pile, matchesToTake := calculateOptimalMove(gameState)

	fmt.Printf("Computer removes from %d matches from pile %d\n", matchesToTake, pile)

	matchesLeftover := len(gameState.GameBoard[pile-1]) - matchesToTake

	gameState.GameBoard[pile-1] = strings.Repeat("|", matchesLeftover)

	return true
}

func calculateOptimalMove(gameState *board.GameState) (int, int) {
	gameBoardLength := len(gameState.GameBoard)

	pileLengthBoard := make([]int, gameBoardLength)

	for i := 0; i < gameBoardLength; i++ {
		pileLengthBoard[i] = len(gameState.GameBoard[i])
	}

	nimSum := calculateNimSum(pileLengthBoard)

	if nimSum == 0 {

		for i := 0; i < gameBoardLength; i++ {

			if pileLengthBoard[i] > 0 {

				return i + 1, 1
			}
		}
	}

	for i := 0; i < gameBoardLength; i++ {

		targetSize := pileLengthBoard[i] ^ nimSum
		if (targetSize) < pileLengthBoard[i] {

			return i + 1, pileLengthBoard[i] - (targetSize)
		}
	}

	panic("This shouldn't happen")
}

func calculateNimSum(pileLengthBoard []int) int {
	nimSum := 0

	for i := 0; i < len(pileLengthBoard); i++ {
		nimSum = nimSum ^ pileLengthBoard[i]
	}

	return nimSum
}
