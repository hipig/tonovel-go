package services

import (
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"golang.org/x/net/html"
	"log"
	"net/url"
)

func ParseRule(node *html.Node, rule string, t string) (result string) {
	doc, err := htmlquery.Query(node, rule)
	if rule == "" || err != nil || doc == nil {
		return
	}
	switch t {
	case "html":
		result = htmlquery.OutputHTML(doc, false)
	default:
		result = htmlquery.InnerText(doc)
	}
	return
}

func FormatURL(href string, base string) (result string) {
	uri, err := url.Parse(href)
	if href == "" || err != nil {
		return
	}
	baseUri, _ := url.Parse(base)
	if !uri.IsAbs() {
		result = baseUri.ResolveReference(uri).String()
	}else {
		result = uri.String()
	}
	return
}

func NewColly(resp colly.ResponseCallback) *colly.Collector {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnResponse(resp)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	return c
}