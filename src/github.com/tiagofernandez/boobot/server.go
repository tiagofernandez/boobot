package boobot

import (
	"github.com/kataras/iris"
	"log"
)

func Start(port string) {
	iris.Get("/", func(c *iris.Context) {
		c.JSON(200, iris.Map{
			"Message": "👋, 🌎!",
		})
	})
	iris.Get("/boobs", func(c *iris.Context) {
		c.JSON(200, iris.Map{
			"URL": GimmeSomeBoobs(),
		})
	})
	iris.Listen(":" + port)
	log.Println("Boobot is online ( o Y o )")
}
