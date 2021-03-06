package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/samburville/GamesPlayingWithComputers/nim/player"

	"github.com/samburville/GamesPlayingWithComputers/nim/consolehelper"

	"github.com/samburville/GamesPlayingWithComputers/nim/board"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	piles, err := consolehelper.AskPlayerForInt("How many piles do you want your game to have?", reader)

	if err != nil {
		panic(err)
	}

	matches, err := consolehelper.AskPlayerForInt("How many matches do you want each pile to have?", reader)

	if err != nil {
		panic(err)
	}

	player1 := player.HumanPlayer{}

	player2 := player.ComputerPlayer{}

	gameState := board.NewGameState(piles, matches)

	playerCounter := 1

	for !gameState.IsGameover() {
		playerCounter = (playerCounter + 1) % 2

		gameState.PrintBoard()
		fmt.Printf("Player %d it is your turn\n", playerCounter+1)

		for !player1.PlayMove(gameState, reader) {
		}

		if gameState.IsGameover() {
			break
		}

		playerCounter = (playerCounter + 1) % 2

		gameState.PrintBoard()
		fmt.Printf("Player %d it is your turn\n", playerCounter+1)

		player2.PlayMove(gameState, reader)
	}

	fmt.Printf("Player %d you have won! You took the last match!\n", playerCounter+1)
}
