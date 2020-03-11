package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"tonovel/bootstrap"
	"tonovel/http/controllers"
)

func Configure(b *bootstrap.Bootstrapper)  {
	mvc.New(b.Party("/")).Handle(new(controllers.IndexController))
	mvc.New(b.Party("/detail")).Handle(new(controllers.DetailController))
	mvc.New(b.Party("/content")).Handle(new(controllers.ContentController))
}