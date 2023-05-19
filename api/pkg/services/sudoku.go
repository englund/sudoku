package services

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

func (ctx context) GetNewGame() (*Board, error) {
	// TODO: error handling
	return createEmptyGame(9), nil
}

func createEmptyGame(size int) *Board {
	a := make([][]int, size)
	for i := range a {
		a[i] = make([]int, size)
	}
	return &a
}
