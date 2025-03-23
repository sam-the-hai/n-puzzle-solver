package solver

import (
	puzzle "n-puzzle-solver/internal/puzzle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFSSolver(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]int
		size     int
		wantPath bool
	}{
		{
			name: "simple solvable puzzle",
			board: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 5, 8},
			},
			size:     3,
			wantPath: true,
		},
		{
			name: "goal state",
			board: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			size:     3,
			wantPath: true,
		},
		{
			name: "unsolvable puzzle",
			board: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{8, 7, 0},
			},
			size:     3,
			wantPath: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := puzzle.NewPuzzle(tt.size)
			p.Board = tt.board
			p.EmptyPos = findEmptyPosition(tt.board)

			solver := NewDFSSolver()
			assert.Equal(t, tt.wantPath, solver.IsSolvable(p))

			if tt.wantPath {
				solution := solver.Solve(p)
				assert.NotNil(t, solution)
				assert.True(t, len(solution) > 0)
			} else {
				solution := solver.Solve(p)
				assert.Nil(t, solution)
			}
		})
	}
}

func findEmptyPosition(board [][]int) puzzle.Position {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				return puzzle.Position{Row: i, Col: j}
			}
		}
	}
	return puzzle.Position{}
}
