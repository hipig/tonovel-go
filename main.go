package main

import (
	"tonovel/bootstrap"
	"tonovel/http/middlewares"
	"tonovel/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("toNovel", "205270006@qq.com")
	app.Bootstrap()
	app.Configure(middlewares.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
