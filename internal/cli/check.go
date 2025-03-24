package cli

import (
	"fmt"
	"n-puzzle-solver/internal/puzzle"
	"n-puzzle-solver/internal/solver"
	"strconv"
)

func CheckSolvability(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: check <size> <numbers...>\nExample: check 3 1 2 3 4 0 6 7 5 8")
	}

	size, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid size: %w", err)
	}

	if len(args) != size*size+1 {
		return fmt.Errorf("expected %d numbers, got %d", size*size, len(args)-1)
	}

	board := make([][]int, size)
	for i := range size {
		board[i] = make([]int, size)
		for j := range size {
			num, err := strconv.Atoi(args[i*size+j+1])
			if err != nil {
				return fmt.Errorf("invalid number at position [%d][%d]: %w", i, j, err)
			}
			board[i][j] = num
		}
	}

	p, err := puzzle.NewPuzzleFromBoard(board)
	if err != nil {
		return fmt.Errorf("invalid puzzle: %w", err)
	}

	fmt.Println("\nPuzzle to check:")
	fmt.Println(p.String())

	s := solver.NewAStarSolver()
	if s.IsSolvable(p) {
		fmt.Println("This puzzle is solvable")
	} else {
		fmt.Println("This puzzle is not solvable")
	}

	return nil
}
