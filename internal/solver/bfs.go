package solver

import (
	"container/list"
	"n-puzzle-solver/internal/puzzle"
)

type BFSSolver struct{}

func NewBFSSolver() Solver {
	return &BFSSolver{}
}

// Node represents a state in the search tree
type BFSNode struct {
	puzzle *puzzle.Puzzle
	parent *BFSNode
	move   string
}

// Solve implements the Breadth-First Search algorithm
func (s *BFSSolver) Solve(p *puzzle.Puzzle) []string {
	if !s.IsSolvable(p) {
		return nil
	}

	// Initialize queue
	queue := list.New()
	queue.PushBack(&BFSNode{
		puzzle: p,
		parent: nil,
		move:   "",
	})

	// Track visited states
	visited := make(map[string]bool)
	visited[boardToString(p.Board)] = true

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(*BFSNode)

		if current.puzzle.IsGoalState() {
			return reconstructBFSPath(current)
		}

		// Try all possible moves
		for direction := range puzzle.Directions {
			if !current.puzzle.IsValidMove(direction) {
				continue
			}

			newPuzzle := applyMove(current.puzzle, direction)
			boardStr := boardToString(newPuzzle.Board)

			if !visited[boardStr] {
				queue.PushBack(&BFSNode{
					puzzle: newPuzzle,
					parent: current,
					move:   direction,
				})
				visited[boardStr] = true
			}
		}
	}

	return nil
}

// IsSolvable checks if the puzzle is solvable
func (s *BFSSolver) IsSolvable(p *puzzle.Puzzle) bool {
	return isSolvable(p)
}

func reconstructBFSPath(node *BFSNode) []string {
	var path []string
	current := node
	for current.parent != nil {
		path = append([]string{current.move}, path...)
		current = current.parent
	}
	return path
}
