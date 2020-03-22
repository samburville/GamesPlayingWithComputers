package player

import (
	"bufio"

	"github.com/samburville/GamesPlayingWithComputers/nim/board"
)

// Player definition of methods required for a player of Nim
type Player interface {
	PlayMove(gameState *board.GameState, reader *bufio.Reader) bool
}
