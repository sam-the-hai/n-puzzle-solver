package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPuzzle(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantErr bool
	}{
		{
			name:    "valid 3x3 puzzle",
			size:    3,
			wantErr: false,
		},
		{
			name:    "valid 4x4 puzzle",
			size:    4,
			wantErr: false,
		},
		{
			name:    "invalid size 1",
			size:    1,
			wantErr: true,
		},
		{
			name:    "invalid negative size",
			size:    -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle, err := NewPuzzle(tt.size)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, puzzle)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, puzzle)
				assert.Equal(t, tt.size, puzzle.Size)
				assert.Equal(t, tt.size, len(puzzle.Board))
				assert.Equal(t, tt.size, len(puzzle.Board[0]))
			}
		})
	}
}

func TestPuzzle_IsGoalState(t *testing.T) {
	puzzle, _ := NewPuzzle(3)
	assert.True(t, puzzle.IsGoalState(), "New puzzle should be in goal state")

	// Modify puzzle state
	puzzle.Board[0][0], puzzle.Board[0][1] = puzzle.Board[0][1], puzzle.Board[0][0]
	assert.False(t, puzzle.IsGoalState(), "Modified puzzle should not be in goal state")
}

func TestPuzzle_IsValidMove(t *testing.T) {
	puzzle, _ := NewPuzzle(3)

	tests := []struct {
		name      string
		direction string
		want      bool
	}{
		{
			name:      "valid up move",
			direction: "UP",
			want:      true,
		},
		{
			name:      "invalid down move",
			direction: "DOWN",
			want:      false,
		},
		{
			name:      "invalid direction",
			direction: "INVALID",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := puzzle.IsValidMove(tt.direction)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPuzzle_Shuffle(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		moves    int
		wantDiff bool
	}{
		{
			name:     "no moves should not change puzzle",
			size:     3,
			moves:    0,
			wantDiff: false,
		},
		{
			name:     "should shuffle puzzle with moves",
			size:     3,
			moves:    20,
			wantDiff: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle, _ := NewPuzzle(tt.size)

			// Store original board state
			original := make([][]int, puzzle.Size)
			for i := range puzzle.Board {
				original[i] = make([]int, puzzle.Size)
				copy(original[i], puzzle.Board[i])
			}

			puzzle.Shuffle(tt.moves)

			// Compare if boards are different
			isDifferent := false
			for i := 0; i < puzzle.Size; i++ {
				for j := 0; j < puzzle.Size; j++ {
					if puzzle.Board[i][j] != original[i][j] {
						isDifferent = true
						break
					}
				}
			}

			assert.Equal(t, tt.wantDiff, isDifferent)

			// Verify that all numbers are still present
			numbers := make(map[int]bool)
			for i := 0; i < puzzle.Size; i++ {
				for j := 0; j < puzzle.Size; j++ {
					numbers[puzzle.Board[i][j]] = true
				}
			}
			assert.Equal(t, puzzle.Size*puzzle.Size, len(numbers))
		})
	}
}
