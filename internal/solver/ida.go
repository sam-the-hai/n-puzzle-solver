package solver

import (
	"n-puzzle-solver/internal/puzzle"
)

type IDASolver struct{}

func NewIDASolver() Solver {
	return &IDASolver{}
}

type IDANode struct {
	puzzle    *puzzle.Puzzle
	parent    *IDANode
	move      string
	cost      int
	heuristic int
}

func (s *IDASolver) Solve(p *puzzle.Puzzle) []string {
	if !s.IsSolvable(p) {
		return nil
	}

	threshold := float64(calculateManhattanDistance(p))
	for {
		result, newThreshold := s.search(p, threshold)
		if result != nil {
			return reconstructIDAPath(result)
		}
		if newThreshold == float64(1<<31-1) {
			return nil
		}
		threshold = newThreshold
	}
}

func (s *IDASolver) search(p *puzzle.Puzzle, threshold float64) (*IDANode, float64) {
	initialNode := &IDANode{
		puzzle:    p,
		parent:    nil,
		move:      "",
		cost:      0,
		heuristic: calculateManhattanDistance(p),
	}

	minCost := float64(1<<31 - 1)
	stack := []*IDANode{initialNode}
	visited := make(map[string]bool)
	visited[boardToString(p.Board)] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		f := float64(current.cost + current.heuristic)
		if f > threshold {
			if f < minCost {
				minCost = f
			}
			continue
		}

		if current.puzzle.IsGoalState() {
			return current, 0
		}

		for direction := range puzzle.Directions {
			if !current.puzzle.IsValidMove(direction) {
				continue
			}

			newPuzzle := applyMove(current.puzzle, direction)
			boardStr := boardToString(newPuzzle.Board)

			if !visited[boardStr] {
				newNode := &IDANode{
					puzzle:    newPuzzle,
					parent:    current,
					move:      direction,
					cost:      current.cost + 1,
					heuristic: calculateManhattanDistance(newPuzzle),
				}
				stack = append(stack, newNode)
				visited[boardStr] = true
			}
		}
	}

	return nil, minCost
}

func (s *IDASolver) IsSolvable(p *puzzle.Puzzle) bool {
	return isSolvable(p)
}

func reconstructIDAPath(node *IDANode) []string {
	var path []string
	current := node
	for current.parent != nil {
		path = append([]string{current.move}, path...)
		current = current.parent
	}
	return path
}
