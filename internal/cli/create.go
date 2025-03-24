package cli

import (
	"fmt"
	"n-puzzle-solver/internal/puzzle"
	"n-puzzle-solver/internal/solver"
	"strconv"
)

func CreateSolvablePuzzle(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: create <size>\nExample: create 3")
	}

	size, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid size: %w", err)
	}

	if size < 3 || size > 5 {
		return fmt.Errorf("size must be between 3 and 5")
	}

	// Create a new puzzle and shuffle until it's solvable
	var p *puzzle.Puzzle
	for {
		p, err = puzzle.NewPuzzle(size)
		if err != nil {
			return fmt.Errorf("failed to create puzzle: %w", err)
		}
		p.Shuffle(20) // Shuffle 20 times for a good mix

		s := solver.NewAStarSolver()
		if s.IsSolvable(p) {
			break
		}
	}

	fmt.Println("\nCreated solvable puzzle:")
	fmt.Println(p.String())

	return nil
}
