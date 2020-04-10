package player

import "testing"

func Test_calculateNimSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"1 pile board case 1", []int{1}, 1},
		{"1 pile board case 2", []int{10}, 10},
		{"2 pile board case 1", []int{1, 1}, 0},
		{"3 pile board case 1", []int{2, 3, 4}, 5},
		{"3 pile board case 2", []int{3, 4, 5}, 2},
		{"3 pile board case 3", []int{1, 0, 8}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateNimSum(tt.input); got != tt.want {
				t.Errorf("calculateNimSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
