package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/go/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		message := "Hello " + name
		c.String(http.StatusOK, message)
	})
	router.GET("/go", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run(":8000")
}
