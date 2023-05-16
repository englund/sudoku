package services

type Board struct {
	Board string
}

type context struct {
}

func NewSudokuService() *context {
	return &context{}
}

func (ctx context) GetNewGame() (*Board, error) {
	// TODO: error handling
	return &Board{Board: "the board"}, nil
}
