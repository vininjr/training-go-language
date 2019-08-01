package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/movies", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNoContent)
	})

	v1.POST("/movies", func(c *gin.Context) {

	})

	router.Run()
}
