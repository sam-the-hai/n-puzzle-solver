# N-Puzzle Solver

A command-line tool for solving the n-puzzle problem using various algorithms. The n-puzzle is a sliding puzzle that consists of a frame of numbered square tiles in random order with one tile missing.

## Features

- Create solvable puzzles of different sizes (3x3 to 5x5)
- Check if a puzzle is solvable
- Solve puzzles using multiple algorithms:
  - A* (default)
  - Breadth First Search (BFS)
  - Depth First Search (DFS)
  - Iterative Deepening A* (IDA*)
  - Greedy Best First

## Installation

```bash
git clone https://github.com/yourusername/n-puzzle-solver.git
cd n-puzzle-solver
go mod download
```

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run benchmarks:
```bash
go test -bench=. ./internal/solver/...
```

## Usage

### Create a Solvable Puzzle

Create a random solvable puzzle of specified size (3-5):
```bash
go run cmd/main.go create 3
```

### Check Puzzle Solvability

Check if a given puzzle is solvable:
```bash
go run cmd/main.go check 3 1 2 3 4 0 6 7 5 8
```

### Solve a Puzzle

Solve a puzzle using a specific algorithm:
```bash
# Using A* (default)
go run cmd/main.go solve 3 1 2 3 4 0 6 7 5 8

# Using BFS
go run cmd/main.go solve -a bfs 3 1 2 3 4 0 6 7 5 8

# Using DFS
go run cmd/main.go solve -a dfs 3 1 2 3 4 0 6 7 5 8

# Using IDA*
go run cmd/main.go solve -a ida 3 1 2 3 4 0 6 7 5 8

# Using Greedy
go run cmd/main.go solve -a greedy 3 1 2 3 4 0 6 7 5 8
```

## Input Format

For the `check` and `solve` commands, input the puzzle as space-separated numbers in row-major order:
- First number: puzzle size (e.g., 3 for 3x3)
- Following numbers: puzzle tiles (0 represents empty tile)

Example for a 3x3 puzzle:
```
1 2 3
4 0 6
7 5 8
```
Should be input as: `3 1 2 3 4 0 6 7 5 8`

## Project Structure

```
n-puzzle-solver/
├── cmd/
│   └── main.go          # Main program entry point
├── internal/
│   ├── cli/             # Command-line interface
│   │   ├── create.go    # Create command
│   │   ├── check.go     # Check command
│   │   └── solve.go     # Solve command
│   ├── puzzle/          # Puzzle domain logic
│   └── solver/          # Solving algorithms
└── README.md
```

## License

MIT License 