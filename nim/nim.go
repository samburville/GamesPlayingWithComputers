package main

import (
	"bufio"
	"errors"
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

	piles := AskPlayerForInt("How many piles do you want your game to have?", reader)

	matches := AskPlayerForInt("How many matches do you want each pile to have?", reader)

	gameState := NewGameState(piles, matches)

	playerCounter := 0

	for !gameState.gameover {
		gameState.printBoard()

		fmt.Printf("Player %d it is your turn\n", playerCounter+1)
		gameState.playMove(reader)

		gameState.gameover = gameState.isGameover()

		playerCounter = (playerCounter + 1) % 2
	}

	gameState.printBoard()
	fmt.Printf("Player %d you have lost! There is only 1 match left\n", playerCounter+1)
}

func (gameState *GameState) playMove(reader *bufio.Reader) {

	pile := AskPlayerForInt("Which pile would you like to remove from?", reader)

	if pile > len(gameState.gameBoard) {
		panic(errors.New("Pile is out of range of the game board"))
	}

	matches := AskPlayerForInt("How many matches would you like to remove?", reader)

	matchesLeftover := len(gameState.gameBoard[pile-1]) - matches

	if matchesLeftover < 0 {
		panic(errors.New("Number of matches to remove is larger than the pile"))
	}

	if matches == gameState.getNumberOfMatchesRemaining() {
		panic(errors.New("Player cannot remove the final match"))
	}

	gameState.gameBoard[pile-1] = strings.Repeat("|", matchesLeftover)
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
func AskPlayerForInt(question string, reader *bufio.Reader) int {
	fmt.Println(question)
	userInput, _ := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\n", "", -1)

	userInputAsInteger, err := strconv.Atoi(userInput)

	if err != nil {
		panic(err)
	}

	fmt.Println()

	return userInputAsInteger
}
