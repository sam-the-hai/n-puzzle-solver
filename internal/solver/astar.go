package solver

import (
	"fmt"
	"n-puzzle-solver/internal/puzzle"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type Solver interface {
	Solve(puzzle *puzzle.Puzzle) []string
	IsSolvable(puzzle *puzzle.Puzzle) bool
}

type AStarSolver struct{}

func NewAStarSolver() Solver {
	return &AStarSolver{}
}

// Node represents a state in the search tree
type Node struct {
	puzzle    *puzzle.Puzzle
	parent    *Node
	move      string
	cost      int
	heuristic int
	index     int
}

// Solve implements the A* algorithm to find the solution
func (s *AStarSolver) Solve(p *puzzle.Puzzle) []string {
	if !s.IsSolvable(p) {
		return nil
	}

	// Create priority queue with custom comparator
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		nodeA := a.(*Node)
		nodeB := b.(*Node)
		priorityA := nodeA.cost + nodeA.heuristic
		priorityB := nodeB.cost + nodeB.heuristic
		return priorityA - priorityB
	})

	initialNode := &Node{
		puzzle:    p,
		parent:    nil,
		move:      "",
		cost:      0,
		heuristic: calculateManhattanDistance(p),
	}

	pq.Enqueue(initialNode)

	visited := make(map[string]bool)
	visited[boardToString(p.Board)] = true

	for !pq.Empty() {
		currentInterface, _ := pq.Dequeue()
		current := currentInterface.(*Node)

		if current.puzzle.IsGoalState() {
			return reconstructPath(current)
		}

		for direction := range puzzle.Directions {
			if !current.puzzle.IsValidMove(direction) {
				continue
			}

			newPuzzle := applyMove(current.puzzle, direction)
			boardStr := boardToString(newPuzzle.Board)

			if !visited[boardStr] {
				newNode := &Node{
					puzzle:    newPuzzle,
					parent:    current,
					move:      direction,
					cost:      current.cost + 1,
					heuristic: calculateManhattanDistance(newPuzzle),
				}
				pq.Enqueue(newNode)
				visited[boardStr] = true
			}
		}
	}

	return nil
}

// IsSolvable checks if the puzzle is solvable
func (s *AStarSolver) IsSolvable(puzzle *puzzle.Puzzle) bool {
	return isSolvable(puzzle)
}

// Helper functions
func calculateManhattanDistance(puzzle *puzzle.Puzzle) int {
	distance := 0
	for i := range puzzle.Size {
		for j := range puzzle.Size {
			if puzzle.Board[i][j] != 0 {
				value := puzzle.Board[i][j]
				targetRow := (value - 1) / puzzle.Size
				targetCol := (value - 1) % puzzle.Size
				distance += abs(i-targetRow) + abs(j-targetCol)
			}
		}
	}
	return distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func applyMove(p *puzzle.Puzzle, direction string) *puzzle.Puzzle {
	newPuzzle, _ := puzzle.NewPuzzle(p.Size)

	// Copy the board
	for i := range p.Board {
		copy(newPuzzle.Board[i], p.Board[i])
	}
	newPuzzle.EmptyPos = p.EmptyPos

	// Get the direction offsets
	dir := puzzle.Directions[direction]
	newRow := p.EmptyPos.Row + dir[0]
	newCol := p.EmptyPos.Col + dir[1]

	// Swap tiles
	newPuzzle.Board[p.EmptyPos.Row][p.EmptyPos.Col] = newPuzzle.Board[newRow][newCol]
	newPuzzle.Board[newRow][newCol] = 0
	newPuzzle.EmptyPos = puzzle.Position{Row: newRow, Col: newCol}

	return newPuzzle
}

func boardToString(board [][]int) string {
	result := ""
	for _, row := range board {
		for _, val := range row {
			result += fmt.Sprintf("%d", val)
		}
	}
	return result
}

func reconstructPath(node *Node) []string {
	var path []string
	current := node
	for current.parent != nil {
		path = append([]string{current.move}, path...)
		current = current.parent
	}
	return path
}
