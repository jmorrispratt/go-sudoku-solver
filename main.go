/*
Application to solve the famous Sudoku puzzle
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// testing with project euler's sudoku set
	projectEulerTester()
}

//projectEulerTester performs tests in all sudokus from the test file (obtained from project euler)
func projectEulerTester() {
	// test cases path
	testCasesPath := "./test_cases.txt"

	// opening the file with the test cases
	f, err := os.Open(testCasesPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	input := bufio.NewScanner(f)

	// there are 50 test cases
	for i := 0; i < 50; i++ {
		// loading sudoku puzzle board
		board := loadTestCase(input)

		// solving sudoku puzzle (naive backtrack implementation)
		answer := SolveSudokuPuzzle(board)

		// problems solving test case i-th
		if answer == nil {
			fmt.Printf("Problems solving puzzle %v\n", i+1)
		} else {
			// notification that i-th puzzle was solved
			fmt.Printf("Puzzle %v solved\n", i+1)
		}
	}

	// closing file
	f.Close()
}

//loadTestCase loads a single sudoku test case from the test cases files
func loadTestCase(input *bufio.Scanner) []int {
	// reading header (i.e. Grid X) and discarding it
	if !input.Scan() {
		fmt.Println("Error in the test cases file format.")
	}
	input.Text()

	// those are 9x9 sudoku puzzles (read 9 lines)
	s := ""
	for i := 0; i < 9; i++ {
		if !input.Scan() {
			fmt.Println("Error in the test cases file format.")
		}

		// doing this by debugging purposes (s += input.Text() would suffice)
		line := input.Text()
		s += line
	}

	return strToIntList(s)
}

//strToIntList converts a string of digits into a []int with such digits
func strToIntList(s string) []int {
	var result = make([]int, len(s))

	// for each character in the string
	for i, c := range s {
		iValue, err := strconv.Atoi(string(c))

		// error handling if int conversion failed
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// setting value in array
		result[i] = iValue
	}

	return result
}
