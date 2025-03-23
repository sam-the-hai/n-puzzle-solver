package solver

import (
	puzzle "n-puzzle-solver/internal/puzzle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFSSolver(t *testing.T) {
	// Create a simple solvable puzzle
	p, _ := puzzle.NewPuzzle(3)
	p.Board = [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}
	p.EmptyPos = puzzle.Position{Row: 1, Col: 1}

	solver := NewBFSSolver()

	// Test IsSolvable
	assert.True(t, solver.IsSolvable(p))

	// Test Solve
	solution := solver.Solve(p)
	assert.NotNil(t, solution)
	assert.True(t, len(solution) > 0)
}

// Test that BFSSolver implements the Solver interface
func TestBFSSolverImplementsInterface(t *testing.T) {
	var _ Solver = (*BFSSolver)(nil)
}
