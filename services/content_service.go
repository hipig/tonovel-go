package services

import (
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
	"tonovel/datamodels"
	"tonovel/datasource"
	"tonovel/repositories"
)

type ContentService interface {
	GetSourceByKey(key string) (datamodels.BookSource, bool)
	GetContent(url string, source string) datamodels.BookContent
}

func NewContentService() ContentService {
	return &contentService{
		rep: repositories.NewBookSourceRepository(datasource.BookSources),
	}
}

type contentService struct {
	rep repositories.BookSourceRepository
}

func (s *contentService) GetContent(url string, key string) (content datamodels.BookContent) {
	source, ok := s.GetSourceByKey(key)
	if !ok {
		content = datamodels.BookContent{}
		return
	}

	NewColly(func(resp *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(resp.Body)))
		if err != nil {
			log.Fatalln(err)
		}

		content = datamodels.BookContent{
			Title:	ParseRule(doc, source.ContentTitleRule, "text"),
			Text:	ParseRule(doc, source.ContentTextRule, "html"),
			PreviousURL:	FormatURL(ParseRule(doc, source.ContentPreviousURLRule, "text"), source.SourceURL),
			NextURL:	FormatURL(ParseRule(doc, source.ContentNextURLRule, "text"), source.SourceURL),
		}
	}).Visit(url)

	return
}

func (s *contentService) GetSourceByKey(key string) (datamodels.BookSource, bool) {
	return s.rep.Select(func(s datamodels.BookSource) bool {
		return s.SourceKey == key
	})
}