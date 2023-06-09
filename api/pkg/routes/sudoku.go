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
	SolveGame(*services.Board) (bool, *services.Board)
}

func Sudoku(g *gin.RouterGroup, ss sudokuService) {
	r := newSudokuRoutes(ss)
	g.GET("/", r.getNewGame)
	g.POST("/solve", r.solveGame)
}

type context struct {
	ss sudokuService
}

func newSudokuRoutes(ss sudokuService) *context {
	return &context{ss}
}

type postSolveGameRequest struct {
	Board *services.Board `json:"board"`
}

type postSolveGameResponse struct {
	Board    *services.Board `json:"board"`
	IsSolved bool            `json:"isSolved"`
}

func (ctx context) solveGame(gCtx *gin.Context) {
	var request postSolveGameRequest
	if gCtx.Bind(&request) != nil {
		gCtx.JSON(http.StatusBadRequest, nil)
		return
	}

	isSolved, board := ctx.ss.SolveGame(request.Board)
	gCtx.JSON(http.StatusOK, mapToSolveGameResponse(board, isSolved))
}

func mapToSolveGameResponse(board *services.Board, isSolved bool) *postSolveGameResponse {
	return &postSolveGameResponse{
		Board:    board,
		IsSolved: isSolved,
	}
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
	gCtx.JSON(http.StatusOK, mapToNewGameResponse(board))
}

func mapToNewGameResponse(board *services.Board) *getNewGameResponse {
	return &getNewGameResponse{
		Board: board,
	}
}
