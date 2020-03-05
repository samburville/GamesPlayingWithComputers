package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GameState structure for playing a game of nim
type GameState struct {
	gameBoard []string
	gameover  bool
}

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

	gameState := NewGameState(piles, matches)

	playerCounter := 0

	for !gameState.gameover {
		gameState.printBoard()

		fmt.Printf("Player %d it is your turn\n", playerCounter+1)

		for !gameState.playMove(reader) {
		}

		gameState.gameover = gameState.isGameover()

		playerCounter = (playerCounter + 1) % 2
	}

	gameState.printBoard()
	fmt.Printf("Player %d you have lost! There is only 1 match left\n", playerCounter+1)
}

func (gameState *GameState) playMove(reader *bufio.Reader) bool {

	pile, err := AskPlayerForInt("Which pile would you like to remove from?", reader)

	if err != nil {
		fmt.Println("Invalid User input. Try Again")
		return false
	}

	if pile > len(gameState.gameBoard) {
		fmt.Println("Pile is out of range of the game board. Try Again")
		fmt.Println()
		return false
	}

	matches, err := AskPlayerForInt("How many matches would you like to remove?", reader)

	if err != nil {
		fmt.Println("Invalid User input. Try Again")
		return false
	}

	matchesLeftover := len(gameState.gameBoard[pile-1]) - matches

	if matchesLeftover < 0 {
		fmt.Println("Number of matches to remove is larger than the pile. Try Again")
		fmt.Println()
		return false
	}

	if matches == gameState.getNumberOfMatchesRemaining() {
		fmt.Println("Player cannot remove the final match. Try Again")
		fmt.Println()
		return false
	}

	gameState.gameBoard[pile-1] = strings.Repeat("|", matchesLeftover)
	return true
}

func (gameState *GameState) printBoard() {
	fmt.Println("Current Game Board")
	fmt.Println()

	for i := range gameState.gameBoard {
		fmt.Println(gameState.gameBoard[i])
	}

	fmt.Println()
}

func (gameState *GameState) isGameover() bool {
	return 1 == gameState.getNumberOfMatchesRemaining()
}

func (gameState *GameState) getNumberOfMatchesRemaining() int {
	matchesRemaining := 0

	for i := range gameState.gameBoard {
		matchesRemaining = matchesRemaining + len(gameState.gameBoard[i])
	}
	return matchesRemaining
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

// AskPlayerForInt ask a player a question and expect an integer response
func AskPlayerForInt(question string, reader *bufio.Reader) (int, error) {
	fmt.Println(question)
	userInput, _ := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\n", "", -1)

	fmt.Println()

	return strconv.Atoi(userInput)
}
