package services

import (
	"errors"
	"math/rand"
	"time"
)

type Board = [][]int

const N = 9

type context struct {
}

func NewSudokuService() *context {
	return &context{}
}

// use naive algorithm to solve the sudoku
func solve(board Board, row, col int) bool {
	if row == N-1 && col == N {
		// end of board, stop
		return true
	}

	if col == N {
		// end of column, go to next row
		row += 1
		col = 0
	}

	if board[row][col] > 0 {
		return solve(board, row, col+1)
	}

	// try all values 1-9
	for val := 1; val <= N; val++ {
		if isCellValid(board, row, col, val) {
			board[row][col] = val
			if solve(board, row, col+1) {
				return true
			}
		}
		board[row][col] = 0
	}

	return false
}

func isCellValid(board Board, row int, col int, val int) bool {
	// check row
	for x := 0; x <= N-1; x++ {
		if board[row][x] == val {
			return false
		}
	}

	// check column
	for x := 0; x <= N-1; x++ {
		if board[x][col] == val {
			return false
		}
	}

	// check block
	blockRow := row - row%3
	blockCol := col - col%3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[blockRow+i][blockCol+j] == val {
				return false
			}
		}
	}

	return true
}

func (ctx context) SolveGame(board *Board) (bool, *Board) {
	isSolved := solve(*board, 0, 0)
	return isSolved, board
}

func getUniqueRandomSlice(size int) []int {
	rand.Seed(time.Now().Unix())
	slice := make([]int, 0)

	// create list of values
	for i := 0; i < size; i++ {
		slice = append(slice, i+1)
	}

	// shuffle
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}

func fillFirstBlockWithRandomValues(board Board) Board {
	rand.Seed(time.Now().Unix())

	randVals := getUniqueRandomSlice(9)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j], randVals = randVals[0], randVals[1:]
		}
	}

	return board
}

func keepRandomNbrOfValues(board Board, keep int) Board {
	for i := 0; i < (N*N)-keep; i++ {
		for {
			row := rand.Intn(9)
			col := rand.Intn(9)
			if board[row][col] != 0 {
				board[row][col] = 0
				break
			}
		}
	}

	return board
}

func (ctx context) GetNewGame() (*Board, error) {
	board := *createEmptyBoard(9)
	board = fillFirstBlockWithRandomValues(board)

	isSolved := solve(board, 0, 0)
	if !isSolved {
		return nil, errors.New("could not create solvable board") // TODO: create api error, or bruteforce retry until found one
	}

	board = keepRandomNbrOfValues(board, 25)

	return &board, nil
}

func createEmptyBoard(size int) *Board {
	a := make([][]int, size)
	for i := range a {
		a[i] = make([]int, size)
	}
	return &a
}
