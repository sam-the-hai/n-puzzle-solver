package solver

import (
	"container/heap"
	"n-puzzle-solver/internal/puzzle"
)

type GreedySolver struct{}

func NewGreedySolver() Solver {
	return &GreedySolver{}
}

type GreedyNode struct {
	puzzle    *puzzle.Puzzle
	parent    *GreedyNode
	move      string
	heuristic int
	index     int
}

type GreedyQueue []*GreedyNode

func (q GreedyQueue) Len() int { return len(q) }

func (q GreedyQueue) Less(i, j int) bool {
	return q[i].heuristic < q[j].heuristic
}

func (q GreedyQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *GreedyQueue) Push(x interface{}) {
	n := len(*q)
	node := x.(*GreedyNode)
	node.index = n
	*q = append(*q, node)
}

func (q *GreedyQueue) Pop() interface{} {
	old := *q
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*q = old[0 : n-1]
	return node
}

func (s *GreedySolver) Solve(p *puzzle.Puzzle) []string {
	if !s.IsSolvable(p) {
		return nil
	}

	// If already at goal state, return empty path
	if p.IsGoalState() {
		return []string{}
	}

	pq := make(GreedyQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &GreedyNode{
		puzzle:    p,
		parent:    nil,
		move:      "",
		heuristic: calculateManhattanDistance(p),
	})

	visited := make(map[string]bool)
	visited[boardToString(p.Board)] = true

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*GreedyNode)

		if current.puzzle.IsGoalState() {
			return reconstructGreedyPath(current)
		}

		for _, direction := range current.puzzle.GetValidMoves() {
			newPuzzle := applyMove(current.puzzle, direction)
			boardStr := boardToString(newPuzzle.Board)

			if !visited[boardStr] {
				heap.Push(&pq, &GreedyNode{
					puzzle:    newPuzzle,
					parent:    current,
					move:      direction,
					heuristic: calculateManhattanDistance(newPuzzle),
				})
				visited[boardStr] = true
			}
		}
	}

	return nil
}

func (s *GreedySolver) IsSolvable(puzzle *puzzle.Puzzle) bool {
	return isSolvable(puzzle)
}

func reconstructGreedyPath(node *GreedyNode) []string {
	var path []string
	current := node
	for current.parent != nil {
		path = append([]string{current.move}, path...)
		current = current.parent
	}
	return path
}
