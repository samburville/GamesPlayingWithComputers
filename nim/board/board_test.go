package board

import (
	"reflect"
	"testing"
)

func TestGameState_GetGameBoard(t *testing.T) {
	tests := []struct {
		name      string
		gameBoard []int
		want      []int
	}{
		{"Simple Get test", []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameState := &GameState{
				gameBoard: tt.gameBoard,
			}
			if got := gameState.GetGameBoard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameState.GetGameBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameState_IsGameover(t *testing.T) {
	tests := []struct {
		name      string
		gameBoard []int
		want      bool
	}{
		{"Gameover 1", []int{}, true},
		{"Gameover 2", []int{0, 0, 0, 0}, true},
		{"Not gameover 1", []int{1, 0, 0}, false},
		{"Not gameover 2", []int{1, 1, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameState := &GameState{
				gameBoard: tt.gameBoard,
			}
			if got := gameState.IsGameover(); got != tt.want {
				t.Errorf("GameState.IsGameover() = %v, want %v", got, tt.want)
			}
		})
	}
}
