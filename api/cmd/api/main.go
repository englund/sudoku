package main

import (
	"net/http"
	"sudoku/api/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Status(r.Group("/status"))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
