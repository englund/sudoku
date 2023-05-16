package routes

import (
	"net/http"
	"sudoku/api/pkg/services"

	"github.com/gin-gonic/gin"
)

type sudokuService interface {
	GetNewGame() (*services.Board, error)
}

func Sudoku(g *gin.RouterGroup, ss sudokuService) {
	r := newSudokuRoutes(ss)
	g.GET("/", r.getNewGame)
}

type context struct {
	ss sudokuService
}

func newSudokuRoutes(ss sudokuService) *context {
	return &context{ss}
}

type getNewGameResponse struct {
	Board string `json:"board"`
}

func (ctx context) getNewGame(gCtx *gin.Context) {
	board, err := ctx.ss.GetNewGame()
	// TODO: error handling
	if err != nil {
		gCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	gCtx.JSON(http.StatusOK, mapToGetNewGameResponse(board))
}

func mapToGetNewGameResponse(board *services.Board) *getNewGameResponse {
	return &getNewGameResponse{
		Board: board.Board,
	}
}
