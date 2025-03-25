//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"n-puzzle-solver/internal/wasm"
)

func main() {
	c := make(chan struct{}, 0)

	// Get the document object
	document := js.Global().Get("document")
	resultDiv := document.Call("getElementById", "result")

	// Generate puzzle function
	js.Global().Set("generatePuzzle", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		puzzle := wasm.GeneratePuzzle(this, args)
		// Convert puzzle to JSON string for display
		puzzleJSON, _ := json.MarshalIndent(puzzle, "", "  ")
		resultDiv.Set("innerHTML", "<pre>"+string(puzzleJSON)+"</pre>")
		return nil
	}))

	// Solve puzzle function
	js.Global().Set("solvePuzzle", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		solution := wasm.SolvePuzzle(this, args)
		// Convert solution to JSON string for display
		solutionJSON, _ := json.MarshalIndent(solution, "", "  ")
		resultDiv.Set("innerHTML", "<pre>"+string(solutionJSON)+"</pre>")
		return nil
	}))

	<-c
}
