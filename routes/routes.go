package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"tonovel/bootstrap"
	"tonovel/http/controllers"
	"tonovel/services"
)

func Configure(b *bootstrap.Bootstrapper)  {
	root := mvc.New(b)
	root.Register(services.NewBookService())
	root.Handle(new(controllers.IndexController))
}