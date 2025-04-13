function startAnimation() {
    if (!window.initialPuzzleState || !window.solutionMoves) {
        console.error("Missing puzzle state or solution moves");
        return;
    }
    animateSolution(window.solutionMoves, window.initialPuzzleState);
}

function applyMove(puzzle, move) {
    const flatPuzzleData = puzzle.flat();
    const size = Math.sqrt(flatPuzzleData.length);
    const newPuzzle = [...flatPuzzleData];
    
    const emptyIndex = newPuzzle.indexOf(0);
    const row = Math.floor(emptyIndex / size);
    const col = emptyIndex % size;
    
    let swapIndex = -1;
    
    switch(move) {
        case "UP":
            if (row < size - 1) swapIndex = emptyIndex + size;
            break;
        case "DOWN":
            if (row > 0) swapIndex = emptyIndex - size;
            break;
        case "LEFT":
            if (col < size - 1) swapIndex = emptyIndex + 1;
            break;
        case "RIGHT":
            if (col > 0) swapIndex = emptyIndex - 1;
            break;
    }
    
    if (swapIndex !== -1) {
        [newPuzzle[emptyIndex], newPuzzle[swapIndex]] = 
        [newPuzzle[swapIndex], newPuzzle[emptyIndex]];
    }
    
    return newPuzzle;
}

function animateSolution(moves, initialPuzzle) {
    console.log("Starting animation with:", { moves, initialPuzzle });
    let currentPuzzle = [...initialPuzzle];
    let currentMove = 0;
    
    // Show initial state
    createPuzzleGrid(currentPuzzle);
    
    const interval = setInterval(() => {
        if (currentMove >= moves.length) {
            clearInterval(interval);
            return;
        }
        
        console.log(`Applying move ${currentMove}: ${moves[currentMove]}`);
        currentPuzzle = applyMove(currentPuzzle, moves[currentMove]);
        createPuzzleGrid(currentPuzzle);
        
        currentMove++;
    }, 1000);
}
