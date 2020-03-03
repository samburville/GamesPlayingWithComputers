package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Game structure for playing a game of nim
type Game struct {
	difficulty     string
	piles, matches int
}

// GameState structure for playing a game of nim
type GameState struct {
	gameBoard []string
	gameover  bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many piles do you want your game to have?\n")
	userInputPiles, _ := reader.ReadString('\n')

	userInputPiles = strings.Replace(userInputPiles, "\n", "", -1)

	piles, err := strconv.Atoi(userInputPiles)

	if err != nil {
		panic(err)
	}

	fmt.Print("How many matches do you want each pile to have?\n")
	userInputMatches, _ := reader.ReadString('\n')

	userInputMatches = strings.Replace(userInputMatches, "\n", "", -1)

	matches, err := strconv.Atoi(userInputMatches)

	if err != nil {
		panic(err)
	}

	fmt.Println(NewGameState(piles, matches))
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
