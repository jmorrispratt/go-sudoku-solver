Go Sudoku Solver!
===================

This is a very **elementary** and **naive** Sudoku solver written in Go.

The **elementary** and **naive** characterization of this implementation is due to the strategy used to solve all the puzzles. The strategy used is the classical [Backtracking](https://en.wikipedia.org/wiki/Backtracking) recursion strategy. All test cases were taken from [Project Euler](http://projecteuler.net) and all of them were correctly solved.

> **Note:**

> - This code is a re-write of [another](https://github.com/jmorrispratt/sudoku_solver)    Sudoku solver written in Ruby.

----------

## Expected input

By convention, empty cells in board are represented with  `0`. A Sudoku board represented used this convention is shown next:

--------+-------+--------
| 0 0 3 | 0 2 0 | 6 0 0 | 
| 9 0 0 | 3 0 5 | 0 0 1 | 
| 0 0 1 | 8 0 6 | 4 0 0 | 
--------+-------+--------
| 0 0 8 | 1 0 2 | 9 0 0 | 
| 7 0 0 | 0 0 0 | 0 0 8 | 
| 0 0 6 | 7 0 8 | 2 0 0 | 
--------+-------+--------
| 0 0 2 | 6 0 9 | 5 0 0 | 
| 8 0 0 | 2 0 3 | 0 0 9 | 
| 0 0 5 | 0 1 0 | 3 0 0 | 
--------+-------+--------

This representation can be simply be obtained by calling function `PrintSudokuBoard` in code file `sudoku_utils.go`. The signature of this function can be seen next:

```go
func PrintSudokuBoard(b []int) {
	// CODE HERE
}
```

As can be seen, the internal representation of a Sudoku board used in this implementation is obtained by appending all rows. Thus, a board can be represented as a single chain of integer number in range `0..9`. For example, the previous `board` would be represented as follows:

`[0 0 3 0 2 0 6 0 0 9 0 0 3 0 5 0 0 1 0 0 1 8 0 6 4 0 0 0 0 8 1 0 2 9 0 0 7 0 0 0 0 0 0 0 8 0 0 6 7 0 8 2 0 0 0 0 2 6 0 9 5 0 0 8 0 0 2 0 3 0 0 9 0 0 5 0 1 0 3 0 0]`

----------

## How does it work?
The implementation finds the first empty cell and and fills it with a valid number (i.e., a number that is not present in the same row, the same column or the 3x3 block). Next, the second empty cell is filled in the same way, and so on to the last empty cell. If the procedure gets to a `dead end` (i.e., no valid digit can be placed in an empty cell), the previous move is `undone` (a.k.a `backtracking`). The same old cell is filled with another valid value and the solver moves on again.

The Sudoku is considered solved if the solver is able to fill all cells with valid digits. The operations performed by this strategy can be visualized next:

![Solving Sudokus using backtracking](https://upload.wikimedia.org/wikipedia/commons/8/8c/Sudoku_solved_by_bactracking.gif)

The visualization above is licensed under the [Creative Commons Attribution-Share Alike 3.0 Unported](https://creativecommons.org/licenses/by-sa/3.0/deed.en). Source: [Wikimedia](https://commons.wikimedia.org/wiki/File:Sudoku_solved_by_bactracking.gif)

----------

## How to use it?

```go
import "main"

func main() {
	// creating a sudoku board
	board := []int{0, 0, 3, 0, 2, 0, 6, 0, 0, 9, 0, 0, 3, 0, 5, 0, 0, 1, 0, 0, 1, 8, 0, 6, 4, 0, 0, 0, 0, 8, 1, 0, 2, 9, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 6, 7, 0, 8, 2, 0, 0, 0, 0, 2, 6, 0, 9, 5, 0, 0, 8, 0, 0, 2, 0, 3, 0, 0, 9, 0, 0, 5, 0, 1, 0, 3, 0, 0}

	// solving sudoku puzzle (naive backtrack implementation)
	answer := SolveSudokuPuzzle(board)
}
```

----------

## Project structure

The implementation is divided in 3 (three) different code files: `main.go`, `sudoku_solvers.go` and `sudoku_utils.go`. This project structure was used in order to improve code readability and maintainability.

In `main.go`, the code consists mainly of an experimental setup, which uses 50 test cases obtained from [Project Euler](http://projecteuler.net). These test cases are all stored in file `test_cases.txt`.

In `sudoku_solvers.go`, the code consists mainly of the implementation of a **naive** sudoku solver. If any contributor would like to implement and test other types of Sudoku solvers, their implementation should be placed here. 

In `sudoku_utils.go`, the code consists mainly of shared functionality of use for every implementation of a Sudoku solver. Any additional features needed in future Sudoku solvers should be placed here.

----------

## License

The code is released under the MIT license.