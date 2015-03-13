package main

import (
	"github.com/plimble/ace"
)

func main() {
	a := ace.Default()
	a.GET("/go", func(c *ace.C) {
		c.String(200, "Hello")
	})
	a.GET("/go/:name", func(c *ace.C) {
		name := c.Params.ByName("name")
		c.String(200, "Hello"+name)
	})
	a.Run(":8000")
}
