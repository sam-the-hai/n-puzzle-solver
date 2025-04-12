//go:build js && wasm
// +build js,wasm

package wasm

import (
	"encoding/json"
	"syscall/js"

	"n-puzzle-solver/internal/puzzle"
	"n-puzzle-solver/internal/solver"
)

type PuzzleRequest struct {
	Board     [][]int `json:"board"`
	Algorithm string  `json:"algorithm"`
}

func GeneratePuzzle(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return map[string]interface{}{
			"error": "Invalid number of arguments",
		}
	}

	size := args[0].Int()

	p, err := puzzle.NewPuzzle(size)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	p.Shuffle(100)

	return map[string]interface{}{
		"puzzle": p.Board,
	}
}

func SolvePuzzle(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return map[string]interface{}{
			"error": "Invalid number of arguments",
		}
	}

	var req PuzzleRequest
	if err := json.Unmarshal([]byte(args[0].String()), &req); err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	p, err := puzzle.NewPuzzle(len(req.Board))
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}
	p.Board = req.Board
	p.EmptyPos = findEmptyPosition(req.Board)

	var s solver.Solver
	switch req.Algorithm {
	case "astar":
		s = solver.NewAStarSolver()
	case "bfs":
		s = solver.NewBFSSolver()
	case "dfs":
		s = solver.NewDFSSolver()
	case "ida":
		s = solver.NewIDASolver()
	case "greedy":
		s = solver.NewGreedySolver()
	default:
		return map[string]interface{}{
			"error": "Invalid algorithm",
		}
	}

	solution := s.Solve(p)
	if solution == nil {
		return map[string]interface{}{
			"error": "No solution found",
		}
	}

	return map[string]interface{}{
		"solution": solution,
	}
}

func findEmptyPosition(board [][]int) puzzle.Position {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				return puzzle.Position{Row: i, Col: j}
			}
		}
	}
	return puzzle.Position{}
}
