package solver

import (
	"fmt"
	"n-puzzle-solver/internal/puzzle"
	"testing"
)

func BenchmarkSolvers(b *testing.B) {
	sizes := []int{3, 4}
	solvers := []struct {
		name   string
		solver Solver
	}{
		{"A*", NewAStarSolver()},
		{"BFS", NewBFSSolver()},
		{"DFS", NewDFSSolver()},
		{"IDA*", NewIDASolver()},
		{"Greedy", NewGreedySolver()},
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			for _, s := range solvers {
				b.Run(s.name, func(b *testing.B) {
					p, _ := puzzle.NewPuzzle(size)
					p.Shuffle(20) // Create a random puzzle state

					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						s.solver.Solve(p)
					}
				})
			}
		})
	}
}
