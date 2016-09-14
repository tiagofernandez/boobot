package goatbot

import (
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
	"log"
)

func Start(port string) {
	iris.Set(
		iris.OptionIsDevelopment(true),
		iris.OptionGzip(true),
		iris.OptionCharset("UTF-8"),
	)
	iris.UseTemplate(
		html.New(html.Config{
			Layout: "layout.html",
		}),
	).Directory("./templates", ".html")
	iris.Static("/static", "./static", 1)
	iris.Get("/", root)
	iris.Get("/boobs", boobs)
	iris.Get("/api/boobs", apiBoobs)
	iris.Listen("0.0.0.0:" + port)
	log.Println("GoatBot is online!") // TODO Make it visible on the console
}

func root(ctx *iris.Context) {
	ctx.Render("index.html",
		map[string]interface{}{"Name": "stranger"},
	)
}

func boobs(ctx *iris.Context) {
	ctx.Data(iris.StatusOK, RandomBoobsImage())
	ctx.SetContentType("image")
}

func apiBoobs(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, iris.Map{
		"url": RandomBoobsUrl(),
	})
}
