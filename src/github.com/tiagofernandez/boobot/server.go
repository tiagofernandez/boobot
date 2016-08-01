package goatbot

import (
	"github.com/kataras/iris"
	"log"
)

func Start(port string) {
	iris.Get("/", func(c *iris.Context) {
		c.JSON(iris.StatusOK, iris.Map{
			"Message": "👋, 🌎!",
		})
	})
	iris.Get("/boobs", func(c *iris.Context) {
		c.Data(iris.StatusOK, GimmeSomeBoobs())
		c.SetContentType("image")
	})
	iris.Listen("0.0.0.0:" + port)
	log.Println("GoatBot is online!")
}
