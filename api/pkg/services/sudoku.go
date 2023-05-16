package services

type Board = [][]uint

type context struct {
}

func NewSudokuService() *context {
	return &context{}
}

func (ctx context) GetNewGame() (*Board, error) {
	// TODO: error handling
	return createEmptyGame(9), nil
}

func createEmptyGame(size int) *Board {
	a := make([][]uint, size)
	for i := range a {
		a[i] = make([]uint, size)
	}
	return &a
}
