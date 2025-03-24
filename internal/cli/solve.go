package cli

import (
	"fmt"
	"n-puzzle-solver/internal/puzzle"
	"n-puzzle-solver/internal/solver"
	"strconv"
	"strings"
)

func SolvePuzzleFromInput(args []string, s solver.Solver) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: solve <size> <numbers...>\nExample: solve 3 1 2 3 4 0 6 7 5 8")
	}

	size, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid size: %w", err)
	}

	if len(args) != size*size+1 {
		return fmt.Errorf("expected %d numbers, got %d", size*size, len(args)-1)
	}

	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
		for j := 0; j < size; j++ {
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

	fmt.Println("\nSolving puzzle:")
	fmt.Println(p.String())

	// Get algorithm name
	algoName := "Unknown"
	switch s.(type) {
	case *solver.AStarSolver:
		algoName = "A*"
	case *solver.BFSSolver:
		algoName = "Breadth First Search"
	case *solver.DFSSolver:
		algoName = "Depth First Search"
	case *solver.IDASolver:
		algoName = "Iterative Deepening A*"
	case *solver.GreedySolver:
		algoName = "Greedy Best First"
	}
	fmt.Printf("\nUsing algorithm: %s\n", algoName)

	if !s.IsSolvable(p) {
		fmt.Println("This puzzle is not solvable")
		return nil
	}

	solution := s.Solve(p)
	if solution == nil {
		fmt.Println("No solution found")
		return nil
	}

	fmt.Printf("Solution found in %d moves\n", len(solution))
	fmt.Printf("Moves: %s\n", strings.Join(solution, " → "))

	return nil
}
