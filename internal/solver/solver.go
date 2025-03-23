package solver

import "n-puzzle-solver/internal/puzzle"

// isSolvable checks if the puzzle is solvable
func isSolvable(puzzle *puzzle.Puzzle) bool {
	inversions := 0
	flatBoard := make([]int, 0, puzzle.Size*puzzle.Size)

	// Flatten the board and count inversions
	for i := range puzzle.Size {
		for j := range puzzle.Size {
			if puzzle.Board[i][j] != 0 {
				flatBoard = append(flatBoard, puzzle.Board[i][j])
			}
		}
	}

	for i := range len(flatBoard) - 1 {
		for j := i + 1; j < len(flatBoard); j++ {
			if flatBoard[i] > flatBoard[j] {
				inversions++
			}
		}
	}

	// For odd-sized puzzles, the number of inversions must be even
	if puzzle.Size%2 == 1 {
		return inversions%2 == 0
	}

	// For even-sized puzzles, the number of inversions plus the row number of the blank
	// from the bottom must be odd
	blankRowFromBottom := puzzle.Size - puzzle.EmptyPos.Row
	return (inversions+blankRowFromBottom)%2 == 1
}
