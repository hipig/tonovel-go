package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tonovel/services"
)

type IndexController struct {
	Ctx iris.Context
}

func (c *IndexController) Get() mvc.View {
	return mvc.View{
		Name: "index.html",
		Data:iris.Map{
			"Title": "首页",
		},
	}
}

func (c *IndexController) GetSearch() mvc.View {
	k := c.Ctx.FormValue("k")
	service := services.NewSearchService()
	results := service.SearchByName(k)

	return mvc.View{
		Name: "search.html",
		Data:iris.Map{
			"Title": k+"-搜索",
			"k": k,
			"results": results,
		},
	}
}