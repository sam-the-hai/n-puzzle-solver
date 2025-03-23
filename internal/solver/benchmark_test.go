package solver

import (
	"fmt"
	"n-puzzle-solver/internal/puzzle"
	"testing"
)

func BenchmarkSolverPerformance(b *testing.B) {
	puzzleSizes := []int{3, 4}
	solvingStrategies := []struct {
		name   string
		solver Solver
	}{
		{"AStar", NewAStarSolver()},
		{"BreadthFirstSearch", NewBFSSolver()},
		{"DepthFirstSearch", NewDFSSolver()},
		{"IterativeDeepeningAStar", NewIDASolver()},
		{"GreedyBestFirst", NewGreedySolver()},
	}

	for _, size := range puzzleSizes {
		b.Run(fmt.Sprintf("PuzzleSize%d", size), func(b *testing.B) {
			for _, strategy := range solvingStrategies {
				b.Run(strategy.name, func(b *testing.B) {
					p, _ := puzzle.NewPuzzle(size)
					p.Shuffle(20)

					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						strategy.solver.Solve(p)
					}
				})
			}
		})
	}
}
