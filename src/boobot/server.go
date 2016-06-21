package boobot

import (
	"github.com/kataras/iris"
)

func Start(port string) {
	iris.Get("/", func(c *iris.Context) {
		c.JSON(200, iris.Map{
			"Message": HelloWorld(),
		})
	})
	iris.Listen(":" + port)
}
