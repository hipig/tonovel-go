package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tonovel/services"
)

type IndexController struct {
	Ctx iris.Context
	Service services.NovelService
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
	results := c.Service.GetListByKeyword(k)

	return mvc.View{
		Name: "search.html",
		Data:iris.Map{
			"Title": k+"-搜索",
			"k": k,
			"results": results,
		},
	}
}

func (c *IndexController) GetDetail() mvc.View {
	url := c.Ctx.FormValue("url")
	source := c.Ctx.FormValue("source")
	info, newChapterList, chapterList := c.Service.GetDetail(url, source)

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

func (c *IndexController) GetRead() mvc.View {
	url := c.Ctx.FormValue("url")
	source := c.Ctx.FormValue("source")
	content := c.Service.GetContent(url, source)

	return mvc.View{
		Name: "content.html",
		Data: iris.Map{
			"Title": content.Title+"-内容",
			"content": content,
			"source": source,
		},
	}
}