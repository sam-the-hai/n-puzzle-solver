package solver

import (
	"n-puzzle-solver/internal/puzzle"
)

type DFSSolver struct{}

func NewDFSSolver() Solver {
	return &DFSSolver{}
}

type DFSNode struct {
	puzzle *puzzle.Puzzle
	parent *DFSNode
	move   string
	depth  int
}

func (s *DFSSolver) Solve(p *puzzle.Puzzle) []string {
	if !s.IsSolvable(p) {
		return nil
	}

	// If already at goal state, return empty path
	if p.IsGoalState() {
		return []string{}
	}

	stack := []*DFSNode{{
		puzzle: p,
		parent: nil,
		move:   "",
		depth:  0,
	}}

	visited := make(map[string]bool)
	visited[boardToString(p.Board)] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.puzzle.IsGoalState() {
			return reconstructDFSPath(current)
		}

		for direction := range puzzle.Directions {
			if !current.puzzle.IsValidMove(direction) {
				continue
			}

			newPuzzle := applyMove(current.puzzle, direction)
			boardStr := boardToString(newPuzzle.Board)

			if !visited[boardStr] {
				stack = append(stack, &DFSNode{
					puzzle: newPuzzle,
					parent: current,
					move:   direction,
					depth:  current.depth + 1,
				})
				visited[boardStr] = true
			}
		}
	}

	return nil
}

func (s *DFSSolver) IsSolvable(p *puzzle.Puzzle) bool {
	return isSolvable(p)
}

func reconstructDFSPath(node *DFSNode) []string {
	var path []string
	current := node
	for current.parent != nil {
		path = append([]string{current.move}, path...)
		current = current.parent
	}
	return path
}
