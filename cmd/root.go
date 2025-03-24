package cmd

import (
	"fmt"
	"n-puzzle-solver/internal/cli"
	"n-puzzle-solver/internal/solver"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "n-puzzle-solver",
	Short: "A solver for the n-puzzle problem",
	Long: `A command line tool to solve the n-puzzle problem.
Supports various solving algorithms including A*, BFS, DFS, IDA*, and Greedy Best First.`,
}

func Execute() error {
	return rootCmd.Execute()
}

var createCmd = &cobra.Command{
	Use:   "create <size>",
	Short: "Create a solvable puzzle",
	Long: `Create a solvable puzzle of the specified size.
Example: create 3

Size must be between 3 and 5.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.CreateSolvablePuzzle(args)
	},
}

var checkCmd = &cobra.Command{
	Use:   "check [size] [numbers...]",
	Short: "Check if a puzzle is solvable",
	Long: `Check if a given puzzle is solvable.
Example: check 3 1 2 3 4 0 6 7 5 8`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.CheckSolvability(args)
	},
}

var algorithm string
var solveCmd = &cobra.Command{
	Use:   "solve [size] [numbers...]",
	Short: "Solve a puzzle",
	Long: `Solve a puzzle using various algorithms.
Example: solve 3 1 2 3 4 0 6 7 5 8

Available algorithms:
- astar: A* algorithm (default)
- bfs: Breadth First Search
- dfs: Depth First Search
- ida: Iterative Deepening A*
- greedy: Greedy Best First`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s solver.Solver
		switch algorithm {
		case "astar":
			s = solver.NewAStarSolver()
		case "bfs":
			s = solver.NewBFSSolver()
		case "dfs":
			s = solver.NewDFSSolver()
		case "ida":
			s = solver.NewIDASolver()
		case "greedy":
			s = solver.NewGreedySolver()
		default:
			return fmt.Errorf("unknown algorithm: %s. Use --help to see available algorithms", algorithm)
		}
		return cli.SolvePuzzleFromInput(args, s)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(solveCmd)

	solveCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "astar", "Algorithm to use for solving")
}
