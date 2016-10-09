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
	// TODO Use a router, this is getting outta control!
	iris.Get("/", root)
	iris.Get("/boobs", getBoobsImage)
	iris.Get("/api/boobs", getBoobsUrl)
	iris.Get("/api/events/:group", getEventList)
	iris.Get("/api/events/:group/:action/:user", handleEventAction) // TODO Use POST
	iris.Listen("0.0.0.0:" + port)
	log.Println("GoatBot is online!") // TODO Make it visible on the console
}

func root(ctx *iris.Context) {
	ctx.Render("index.html",
		map[string]interface{}{"Name": "stranger"},
	)
}

func getBoobsImage(ctx *iris.Context) {
	ctx.Data(iris.StatusOK, RandomBoobsImage())
	ctx.SetContentType("image")
}

func getBoobsUrl(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, iris.Map{
		"url": RandomBoobsUrl(),
	})
}

func getEventList(ctx *iris.Context) {
	group := ctx.Param("group")
	ctx.JSON(iris.StatusOK, iris.Map{
		"confirmed": ConfirmedList(group),
	})
}

func handleEventAction(ctx *iris.Context) {
	group := ctx.Param("group")
	user := ctx.Param("user")
	switch ctx.Param("action") {
	case "confirm":
		Confirm(group, user)
	case "cancel":
		Cancel(group, user)
	}
	getEventList(ctx)
}
