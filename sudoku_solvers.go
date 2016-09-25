/*
Solver for the Sudoku puzzle
*/
package main

// --------------------------------------

// sudoku board
var board []int

// board length (i.e. 81)
var boardLength int

// board size (i.e. board width or height)
var boardSize int

// block size (i.e. board block size)
var blockSize int

// possible values (i.e. 1, 2, ..., boardSize)
var allValues []int

// labels for each row
var rowLabels []int

// labels for each column
var colLabels []int

// labels for each internal block
var sqrLabels []int

// --------------------------------------

//SolveSudokuPuzzle is a naive implementation (simple backtracking) of a Sudoku puzzle solver
func SolveSudokuPuzzle(board []int) []int {
	// validating board
	if !isBoardValid(board) {
		return nil
	}

	// computing labels for row, cols, blocks and checking 'correctess'
	rowLabels, colLabels, sqrLabels = getLabels()
	if rowLabels == nil || colLabels == nil || sqrLabels == nil {
		return nil
	}

	// getting board's empty positions
	emptyList := getEmptyPositions(board)

	// returning the solution found via backtrack strategy
	return solveSudokuBacktrack(board, emptyList)
}

//solveSudokuBacktrack finds the solution by means of a classic backtrack strategy
func solveSudokuBacktrack(board []int, emptyList []int) []int {
	// base case: no empty positions to fill
	if len(emptyList) < 1 {
		return board
	}

	// getting first empty index (to attempt to fill correctly)
	index := emptyList[0]

	// getting possible values for this position
	invalidValues := getInvalidValues(board, index)
	possibleValues := getPossibleValues(invalidValues)

	// attempting to solve with each possible value
	for _, j := range possibleValues {
		// setting possible value
		board[index] = j

		// solving by backtraking
		solution := solveSudokuBacktrack(board, emptyList[1:len(emptyList)])

		// if solved then return solution
		if solution != nil {
			return solution
		}

		// taking back that change
		board[index] = 0
	}

	// no solution found
	return nil
}
