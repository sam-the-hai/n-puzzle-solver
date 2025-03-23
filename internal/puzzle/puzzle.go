package puzzle

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Position struct {
	Row int
	Col int
}

type Puzzle struct {
	Size      int
	Board     [][]int
	EmptyPos  Position
	GoalState [][]int
}

var (
	ErrInvalidSize = errors.New("invalid puzzle size: must be greater than 2")
	Directions     = map[string][2]int{
		"UP":    {-1, 0},
		"DOWN":  {1, 0},
		"LEFT":  {0, -1},
		"RIGHT": {0, 1},
	}
)

func NewPuzzle(size int) (*Puzzle, error) {
	if size <= 2 {
		return nil, ErrInvalidSize
	}

	p := &Puzzle{
		Size:      size,
		Board:     make([][]int, size),
		GoalState: make([][]int, size),
	}

	counter := 1
	for i := range size {
		p.Board[i] = make([]int, size)
		p.GoalState[i] = make([]int, size)
		for j := range size {
			if i == size-1 && j == size-1 {
				p.Board[i][j] = 0
				p.GoalState[i][j] = 0
				p.EmptyPos = Position{i, j}
			} else {
				p.Board[i][j] = counter
				p.GoalState[i][j] = counter
				counter++
			}
		}
	}

	return p, nil
}

func (p *Puzzle) IsGoalState() bool {
	for i := range p.Size {
		for j := range p.Size {
			if p.Board[i][j] != p.GoalState[i][j] {
				return false
			}
		}
	}
	return true
}

func (p *Puzzle) IsValidMove(direction string) bool {
	dir, ok := Directions[direction]
	if !ok {
		return false
	}

	newRow := p.EmptyPos.Row + dir[0]
	newCol := p.EmptyPos.Col + dir[1]

	return newRow >= 0 && newRow < p.Size && newCol >= 0 && newCol < p.Size
}

func (p *Puzzle) String() string {
	var result string
	for i := range p.Size {
		for j := range p.Size {
			result += fmt.Sprintf("%2d ", p.Board[i][j])
		}
		result += "\n"
	}
	return result
}

// Shuffle randomly shuffles the puzzle board with the given number of moves
func (p *Puzzle) Shuffle(moves int) {
	if moves <= 0 {
		return
	}

	// Create a new random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for range moves {
		// Get all possible moves
		validMoves := p.GetValidMoves()
		if len(validMoves) == 0 {
			return
		}

		// Choose a random move
		move := validMoves[r.Intn(len(validMoves))]

		// Apply the move
		newRow := p.EmptyPos.Row + Directions[move][0]
		newCol := p.EmptyPos.Col + Directions[move][1]

		// Swap the empty tile with the chosen tile
		p.Board[p.EmptyPos.Row][p.EmptyPos.Col] = p.Board[newRow][newCol]
		p.Board[newRow][newCol] = 0
		p.EmptyPos = Position{Row: newRow, Col: newCol}
	}
}

// GetValidMoves returns a slice of valid move directions
func (p *Puzzle) GetValidMoves() []string {
	var validMoves []string

	for move, dir := range Directions {
		newRow := p.EmptyPos.Row + dir[0]
		newCol := p.EmptyPos.Col + dir[1]

		if newRow >= 0 && newRow < p.Size && newCol >= 0 && newCol < p.Size {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}
