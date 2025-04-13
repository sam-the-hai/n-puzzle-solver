// Initialize WebAssembly
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
    .then((result) => {
        go.run(result.instance);
    })
    .catch(error => {
        console.error('Error loading WebAssembly:', error);
    });

// Event handlers
function handleGenerate() {
    try {
        const size = parseInt(document.getElementById('size').value);
        generatePuzzle(size);
        
        const resultDiv = document.getElementById('result');
        const puzzleData = JSON.parse(resultDiv.textContent);
        
        createPuzzleGrid(puzzleData.puzzle);
    } catch (error) {
        console.error('Error in handleGenerate:', error);
        alert('Error generating puzzle. Please try again.');
    }
}

function handleSolve() {
    try {
        const resultDiv = document.getElementById('result');
        if (!resultDiv || !resultDiv.textContent) {
            alert('Please generate a puzzle first! 🎲');
            return;
        }

        const puzzleData = JSON.parse(resultDiv.textContent);
        window.initialPuzzleState = puzzleData.puzzle;

        const algorithm = document.getElementById('algorithm').value;
        const request = {
            board: puzzleData.puzzle,
            algorithm: algorithm
        };

        solvePuzzle(JSON.stringify(request));
        
        const solutionText = document.getElementById('result').textContent;
        const solutionData = JSON.parse(solutionText.replace(/<\/?pre>/g, ''));
        displaySolution(solutionData);
    } catch (error) {
        console.error('Error in handleSolve:', error);
        alert('Error solving puzzle. Please try again.');
    }
}
