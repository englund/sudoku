package routes

import (
	"log"
	"net/http"
	"sudoku/api/pkg/errors"
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
	Board *services.Board `json:"board"`
}

func (ctx context) getNewGame(gCtx *gin.Context) {
	board, err := ctx.ss.GetNewGame()
	if err != nil {
		switch e := err.(type) {
		case *errors.ApiError: // TODO: may be overkill for this project
			gCtx.JSON(http.StatusInternalServerError, e)
			return
		default:
			log.Println(e)
			gCtx.JSON(http.StatusInternalServerError, errors.NewUnknownApiError(err))
			return
		}
	}
	gCtx.JSON(http.StatusOK, mapToGetNewGameResponse(board))
}

func mapToGetNewGameResponse(board *services.Board) *getNewGameResponse {
	return &getNewGameResponse{
		Board: board,
	}
}
