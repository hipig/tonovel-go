package routes

import (
	"github.com/kataras/iris/v12"
	"tonovel/bootstrap"
)

func Configure(b *bootstrap.Bootstrapper)  {
	b.Get("/", func (ctx iris.Context) {
		ctx.ViewData("Title", "首页")
		ctx.View("index.html")
	})
}