package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tonovel/services"
)

type IndexController struct {
	Ctx iris.Context
	Service services.BookService
}

func (c *IndexController) Get() mvc.View {
	return mvc.View{
		Name: "index.html",
		Data:iris.Map{
			"Title": "首页",
		},
	}
}

func (c *IndexController) GetSearch() {
	k := c.Ctx.FormValue("k")
	results := c.Service.GetListByKeyword(k)

	c.Ctx.JSON(results)
}

func (c *IndexController) GetInfo() {
	detailURL := c.Ctx.FormValue("detail_url")
	source := c.Ctx.FormValue("source")
	info := c.Service.GetInfo(detailURL, source)

	c.Ctx.JSON(info)
}

func (c *IndexController) GetChapters() {
	detailURL := c.Ctx.FormValue("detail_url")
	source := c.Ctx.FormValue("source")
	chapterList := c.Service.GetChapterList(detailURL, source)

	c.Ctx.JSON(chapterList)
}

func (c *IndexController) GetRead() {
	detailURL := c.Ctx.FormValue("detail_url")
	chapterURL := c.Ctx.FormValue("chapter_url")
	source := c.Ctx.FormValue("source")
	content := c.Service.GetContent(detailURL, chapterURL, source)

	c.Ctx.JSON(content)
}
