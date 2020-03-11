package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tonovel/services"
)

type ContentController struct {
	Ctx iris.Context
}

func (c *ContentController) Get() mvc.View {
	url := c.Ctx.FormValue("url")
	source := c.Ctx.FormValue("source")
	service := services.NewContentService()

	content := service.GetContent(url, source)

	return mvc.View{
		Name: "content.html",
		Data: iris.Map{
			"Title": content.Title+"-内容",
			"content": content,
			"source": source,
		},
	}
}