package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"running": true,
		})
	})
}
