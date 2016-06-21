package main

import (
	"github.com/kataras/iris"
	"os"
)

func main() {
	iris.Get("/", func(c *iris.Context) {
		c.JSON(200, iris.Map{
			"Hello": "👋",
			"World": "🌎",
		})
	})
	iris.Listen(":" + os.Args[1])
}
