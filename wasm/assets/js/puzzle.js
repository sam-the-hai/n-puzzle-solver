function createPuzzleGrid(puzzleData) {
    const flatPuzzleData = puzzleData.flat();
    const container = document.getElementById('puzzle-container');
    const size = flatPuzzleData.length;

    const grid = document.createElement('div');
    const linearSize = Math.sqrt(size);
    grid.className = 'puzzle-grid';
    grid.style.gridTemplateColumns = `repeat(${linearSize}, 1fr)`;
    
    flatPuzzleData.forEach((num) => {
        const tile = document.createElement('div');
        tile.className = `puzzle-tile ${num === 0 ? 'empty' : ''}`;
        tile.textContent = num === 0 ? '' : num;
        grid.appendChild(tile);
    });
    
    container.innerHTML = '';
    container.appendChild(grid);
}

function displaySolution(solution) {
    const container = document.getElementById('solution-container');
    if (!solution || !solution.solution) {
        container.innerHTML = '<p>No solution found!</p>';
        return;
    }
    
    window.solutionMoves = solution.solution;
    
    container.innerHTML = `
        <h3>✨ Solution Found!</h3>
        <p class="step-counter">Number of moves: ${solution.solution.length}</p>
        <button onclick="startAnimation()">🎬 Show Solution</button>
    `;
}
