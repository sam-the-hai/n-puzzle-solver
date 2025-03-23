package cli

import (
	"fmt"
	"n-puzzle-solver/internal/puzzle"
	"n-puzzle-solver/internal/solver"

	"github.com/spf13/cobra"
)

func NewApp() *cobra.Command {
	var size int
	var algorithm string

	rootCmd := &cobra.Command{
		Use:   "n-puzzle-solver",
		Short: "A CLI tool to solve n-puzzle problems",
		Long:  `A CLI application that creates and solves n-puzzle problems using various algorithms.`,
	}

	solveCmd := &cobra.Command{
		Use:   "solve",
		Short: "Solve an n-puzzle",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSolve(size, algorithm, cmd)
		},
	}

	solveCmd.Flags().IntVarP(&size, "size", "s", 3, "Size of the puzzle (minimum 3)")
	solveCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "astar", "Solving algorithm (astar, bfs, ida, dfs, or greedy)")
	rootCmd.AddCommand(solveCmd)

	return rootCmd
}

func runSolve(size int, algorithm string, cmd *cobra.Command) error {
	p, err := puzzle.NewPuzzle(size)
	if err != nil {
		return err
	}

	// Shuffle the puzzle
	p.Shuffle(20)

	fmt.Fprintln(cmd.OutOrStdout(), "Initial puzzle state:")
	fmt.Fprintln(cmd.OutOrStdout(), p)

	// Choose solver based on algorithm
	var s solver.Solver
	switch algorithm {
	case "astar":
		s = solver.NewAStarSolver()
	case "bfs":
		s = solver.NewBFSSolver()
	case "ida":
		s = solver.NewIDASolver()
	case "dfs":
		s = solver.NewDFSSolver()
	case "greedy":
		s = solver.NewGreedySolver()
	default:
		return fmt.Errorf("unknown algorithm: %s", algorithm)
	}

	if !s.IsSolvable(p) {
		return fmt.Errorf("puzzle is not solvable")
	}

	solution := s.Solve(p)
	if solution == nil {
		return fmt.Errorf("no solution found")
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Solution found in %d moves:\n", len(solution))
	for i, move := range solution {
		fmt.Fprintf(cmd.OutOrStdout(), "%d. %s\n", i+1, move)
	}

	return nil
}
