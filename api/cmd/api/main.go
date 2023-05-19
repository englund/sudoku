package main

import (
	"sudoku/api/pkg/routes"
	"sudoku/api/pkg/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.Status(r.Group("/status"))

	v1 := r.Group("/v1")
	{
		sudokuService := services.NewSudokuService()
		routes.Sudoku(v1.Group("/sudoku"), sudokuService)
	}

	r.Run(":8080")
}
