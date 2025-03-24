package solver

import (
	puzzle "n-puzzle-solver/internal/puzzle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDASolver(t *testing.T) {
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
		{
			name: "complex puzzle",
			board: [][]int{
				{2, 8, 3},
				{1, 6, 4},
				{7, 0, 5},
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

			// Create and set goal state
			goalState := createGoalState(tt.size)
			p.GoalState = goalState

			solver := NewIDASolver()
			isSolvable := solver.IsSolvable(p)
			t.Logf("Puzzle:\n%v\nIs solvable: %v", p, isSolvable)
			t.Logf("Empty position: %v", p.EmptyPos)
			t.Logf("Goal state:\n%v", p.GoalState)

			// Calculate inversions manually for debugging
			inversions := 0
			flatBoard := make([]int, 0, tt.size*tt.size)
			for i := range tt.size {
				for j := range tt.size {
					if tt.board[i][j] != 0 {
						flatBoard = append(flatBoard, tt.board[i][j])
					}
				}
			}
			for i := range len(flatBoard) - 1 {
				for j := i + 1; j < len(flatBoard); j++ {
					if flatBoard[i] > flatBoard[j] {
						inversions++
					}
				}
			}
			t.Logf("Flat board: %v", flatBoard)
			t.Logf("Inversions: %d", inversions)
			t.Logf("Blank row from bottom: %d", tt.size-p.EmptyPos.Row)

			assert.Equal(t, tt.wantPath, isSolvable)

			if tt.wantPath {
				solution := solver.Solve(p)
				assert.NotNil(t, solution)
				if !p.IsGoalState() {
					assert.True(t, len(solution) > 0)
				}
			} else {
				solution := solver.Solve(p)
				assert.Nil(t, solution)
			}
		})
	}
}

func createGoalState(size int) [][]int {
	goal := make([][]int, size)
	for i := range size {
		goal[i] = make([]int, size)
		for j := range size {
			if i == size-1 && j == size-1 {
				goal[i][j] = 0
			} else {
				goal[i][j] = i*size + j + 1
			}
		}
	}
	return goal
}
