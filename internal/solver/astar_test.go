package solver

import (
	domain "n-puzzle-solver/internal/puzzle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver_Solve(t *testing.T) {
	puzzle, _ := domain.NewPuzzle(3)

	// Make a simple shuffle that we know the solution for
	puzzle.Board = [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}
	puzzle.EmptyPos = domain.Position{Row: 1, Col: 1}

	solver := NewAStarSolver()
	solution := solver.Solve(puzzle)

	assert.NotNil(t, solution)
	assert.True(t, len(solution) > 0)
}

func TestSolver_IsSolvable(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]int
		size     int
		expected bool
	}{
		{
			name: "solvable puzzle",
			board: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 5, 8},
			},
			size:     3,
			expected: true,
		},
		{
			name: "unsolvable puzzle",
			board: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{8, 7, 0},
			},
			size:     3,
			expected: false,
		},
	}

	solver := NewAStarSolver()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle, _ := domain.NewPuzzle(tt.size)
			puzzle.Board = tt.board
			result := solver.IsSolvable(puzzle)
			assert.Equal(t, tt.expected, result)
		})
	}
}
