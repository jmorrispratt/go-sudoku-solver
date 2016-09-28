/*
Application to solve the famous Sudoku puzzle
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	// cli package
	"github.com/urfave/cli"
)

func main() {
	// building and running a cli app
	cliApp := buildCliApp()
	cliApp.Run(os.Args)
}

// buildCliApp configures a cli app
func buildCliApp() (cliApp *cli.App) {
	// creating the app
	app := cli.NewApp()

	// setting app general information
	app.Name = "Sudoku solver"
	app.Usage = "A Go app designed to solve Sudoku puzzles"

	// setting author
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Juan Mauricio Prat",
			Email: "jmorrispratt@gmail.com",
		},
	}

	// setting app version
	app.Version = "0.0.1"

	// setting commands
	app.Commands = []cli.Command{
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Tests the solver against 50 puzzles from Project Euler",
			Action: func(c *cli.Context) error {
				// testing with project euler's sudoku set
				projectEulerTester()
				return nil
			},
		},
	}

	// setting flags (a.k.a options)
	var boardStr string
	var boardPath string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "solve, s",
			Usage:       "Solves Sudoku puzzle with inline board provided",
			Destination: &boardStr,
		},
		cli.StringFlag{
			Name:        "solve-file, f",
			Usage:       "Solves the Sudoku puzzle stored in `FILE`",
			Destination: &boardPath,
		},
	}

	// actions to perform for flags
	app.Action = func(c *cli.Context) error {
		// both choices are not allowed at the same time
		if boardStr != "" && boardPath != "" {
			fmt.Println("Only one of `solve` or `solve-file` is allowed at a time")
		}

		// at least one choice should be selected
		if boardStr == "" && boardPath == "" {
			fmt.Println("At least `solve` or `solve-file` should be selected")
		}

		var answer []int

		// if solving inline puzzle was selected
		if boardStr != "" {
			// building board
			boardInt := strToIntList(boardStr)

			// validating Sudoku board
			if boardInt != nil {
				// solving sudoku puzzle
				answer = SolveSudokuPuzzle(boardInt)
			}
		}

		// solving file-stored puzzle was selected
		if boardPath != "" {
			// opening the file with the test cases
			f, err := os.Open(boardPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			input := bufio.NewScanner(f)

			// building board
			boardPathInt := loadTestCase(input)

			// validating Sudoku board
			if boardPathInt != nil {
				// solving sudoku puzzle
				answer = SolveSudokuPuzzle(boardPathInt)
			}
		}

		// problems solving test case i-th
		if answer == nil {
			fmt.Printf("Problems solving supplied puzzle\n")
		} else {
			// notification that i-th puzzle was solved
			fmt.Printf("Solution:\n")
			PrintSudokuBoard(answer)
		}

		return nil
	}

	// returning the built app
	return app
}

//projectEulerTester performs tests in all sudokus from the test file (obtained from project euler)
func projectEulerTester() {
	// test cases path
	testCasesPath := "./puzzles/test_cases.txt"

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
			fmt.Printf("Puzzle %v (Solution):\n", i+1)
			PrintSudokuBoard(answer)
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
