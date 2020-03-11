package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tonovel/services"
)

type DetailController struct {
	Ctx iris.Context
}

func (c *DetailController) Get() mvc.View {
	url := c.Ctx.FormValue("url")
	source := c.Ctx.FormValue("source")
	service := services.NewDetailService()

	info, newChapterList, chapterList := service.GetDetail(url, source)

	return mvc.View{
		Name: "detail.html",
		Data: iris.Map{
			"Title": info.Name+"-小说详情",
			"url": url,
			"info": info,
			"newChapterList": newChapterList,
			"chapterList": chapterList,
		},
	}
}