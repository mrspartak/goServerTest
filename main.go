package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	router.GET("/go", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run(":8000")
}
