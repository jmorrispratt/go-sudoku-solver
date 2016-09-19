/*
Solver for the Sudoku puzzle
*/
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

//PrintSudokuBoard prints a sudoku board
func PrintSudokuBoard(b []int) {
	bLength := len(b)
	dim := int(math.Sqrt(float64(bLength)))
	blkSize := int(math.Sqrt(float64(dim)))

	blkParts := make([]string, blkSize)
	for i := range blkParts {
		blkParts[i] = strings.Repeat("-", blkSize*2+1)
	}

	hLine := fmt.Sprintf("-%v-", strings.Join(blkParts, "+"))
	for r := 0; r < dim; r++ {
		if r%blkSize == 0 {
			fmt.Println(hLine)
		}

		currLine := "| "
		for c := 0; c < dim; c++ {
			currLine += fmt.Sprintf("%v ", strconv.Itoa(b[r*dim+c]))

			if c%blkSize == (blkSize - 1) {
				currLine += "| "
			}
		}

		fmt.Println(currLine)
	}

	fmt.Println(hLine)
}

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

//getEmptyPositions gets the empty positions in the board
func getEmptyPositions(board []int) []int {
	var result []int

	// getting empy position indexes
	for i := 0; i < boardLength; i++ {
		if board[i] == 0 {
			result = append(result, i)
		}
	}

	return result
}

// isBoardValid checks if the sudoku board is valid
func isBoardValid(board []int) bool {
	// computing and storing board params
	computeBoardParams(board)

	// blockSize must be a  perfect square
	if boardSize*boardSize != boardLength {
		return false
	}

	// blockSize must be again a perfect square number
	if blockSize*blockSize != boardSize {
		return false
	}

	// all good
	return true
}

func getLabels() ([]int, []int, []int) {
	// rows, columns and squares labels for an index on the original board
	rows := make([]int, boardLength)
	cols := make([]int, boardLength)
	squares := make([]int, boardLength)

	// for each index in the original board format
	for i := 0; i < boardLength; i++ {
		rowLabel := i / boardSize
		colLabel := i % boardSize

		a := boardLength / blockSize
		b := rowLabel / blockSize
		c := colLabel / blockSize
		squareLabel := a*b + c

		rows[i] = rowLabel
		cols[i] = colLabel
		squares[i] = squareLabel
	}

	return rows, cols, squares
}

// getInvalidValues gets a list of invalid values for a given board position
func getInvalidValues(board []int, k int) []int {
	var result []int

	// for each item in board
	for j := 0; j < boardLength; j++ {
		// adding all values in the same row, column and block
		if colLabels[k] == colLabels[j] || rowLabels[k] == rowLabels[j] || sqrLabels[k] == sqrLabels[j] {
			// the value at j is not valid
			result = append(result, board[j])
		}
	}

	return result
}

// getPossibleValues gets the possible values given the invalid ones
func getPossibleValues(invalidValues []int) []int {
	return setDiff(allValues, invalidValues)
}

func setDiff(a []int, b []int) []int {
	var diff []int

	// computing set difference by hand (O(n^2) implementation)
	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				found = true
				break
			}
		}
		// if current item was not foud in the other 'set'
		if !found {
			diff = append(diff, a[i])
		}
	}

	return diff
}

// computeBoardParams computes basic params of a sudoku board
func computeBoardParams(board []int) {
	board = board
	boardLength = len(board)
	boardSize = int(math.Sqrt(float64(boardLength)))
	blockSize = int(math.Sqrt(float64(boardSize)))

	// setting all valid values
	allValues = make([]int, boardSize)
	for i := 0; i < boardSize; i++ {
		allValues[i] = i + 1
	}
}
