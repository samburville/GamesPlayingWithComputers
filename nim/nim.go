package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samburville/GamesPlayingWithComputers/nim/board"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	piles, err := AskPlayerForInt("How many piles do you want your game to have?", reader)

	if err != nil {
		panic(err)
	}

	matches, err := AskPlayerForInt("How many matches do you want each pile to have?", reader)

	if err != nil {
		panic(err)
	}

	gameState := board.NewGameState(piles, matches)

	playerCounter := 0

	for !gameState.Gameover {
		gameState.PrintBoard()

		fmt.Printf("Player %d it is your turn\n", playerCounter+1)

		for !playMove(gameState, reader) {
		}

		gameState.Gameover = gameState.IsGameover()

		playerCounter = (playerCounter + 1) % 2
	}

	gameState.PrintBoard()
	fmt.Printf("Player %d you have lost! There is only 1 match left\n", playerCounter+1)
}

func playMove(gameState *board.GameState, reader *bufio.Reader) bool {

	pile, err := AskPlayerForInt("Which pile would you like to remove from?", reader)

	if err != nil {
		fmt.Println("Invalid User input. Try Again")
		return false
	}

	if pile > len(gameState.GameBoard) {
		fmt.Println("Pile is out of range of the game board. Try Again")
		fmt.Println()
		return false
	}

	matches, err := AskPlayerForInt("How many matches would you like to remove?", reader)

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

	if matches == gameState.GetNumberOfMatchesRemaining() {
		fmt.Println("Player cannot remove the final match. Try Again")
		fmt.Println()
		return false
	}

	gameState.GameBoard[pile-1] = strings.Repeat("|", matchesLeftover)
	return true
}

// AskPlayerForInt ask a player a question and expect an integer response
func AskPlayerForInt(question string, reader *bufio.Reader) (int, error) {
	fmt.Println(question)
	userInput, _ := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\n", "", -1)

	fmt.Println()

	return strconv.Atoi(userInput)
}
