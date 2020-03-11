package services

import (
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"golang.org/x/net/html"
	"log"
	"strings"
	"tonovel/datamodels"
	"tonovel/datasource"
	"tonovel/repositories"
)

type DetailService interface {
	GetSourceByKey(key string) (datamodels.BookSource, bool)
	GetDetail(url string, source string) (datamodels.BookInfo, []datamodels.Chapter, []datamodels.Chapter)
	GetDetailInfo(source datamodels.BookSource, doc *html.Node) (info datamodels.BookInfo)
	GetDetailNewChapterList(source datamodels.BookSource, doc *html.Node, url string) (newChapterList []datamodels.Chapter)
	GetDetailChapterList(source datamodels.BookSource, doc *html.Node, url string) (chapterList []datamodels.Chapter)
}

func NewDetailService() DetailService {
	return &detailService{
		repositories.NewBookSourceRepository(datasource.BookSources),
	}
}

type detailService struct {
	rep repositories.BookSourceRepository
}

func (s *detailService) GetDetail(url string, key string) (info datamodels.BookInfo, newChapterList []datamodels.Chapter, chapterList []datamodels.Chapter) {
	source, ok := s.GetSourceByKey(key)
	if !ok {
		info = datamodels.BookInfo{}
		return
	}

	var chapterListURL string
	NewColly(func(resp *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(resp.Body)))
		if err != nil {
			log.Fatalln(err)
		}

		info = s.GetDetailInfo(source, doc)
		if source.DetailChapterListURLRule != "" {
			chapterListURL = ParseRule(doc, source.DetailChapterListURLRule, "text")
		}else {
			newChapterList = s.GetDetailNewChapterList(source, doc, url)
			chapterList = s.GetDetailChapterList(source, doc, url)
			return
		}
	}).Visit(url)

	if chapterListURL != "" {
		NewColly(func(resp *colly.Response) {
			doc, err := htmlquery.Parse(strings.NewReader(string(resp.Body)))
			if err != nil {
				log.Fatalln(err)
			}

			newChapterList = s.GetDetailNewChapterList(source, doc, chapterListURL)
			chapterList = s.GetDetailChapterList(source, doc, chapterListURL)
		}).Visit(chapterListURL)
	}

	return
}

func (s *detailService) GetDetailInfo(source datamodels.BookSource, doc *html.Node) (info datamodels.BookInfo) {
	info = datamodels.BookInfo{
		Name:	ParseRule(doc, source.DetailBookNameRule, "text"),
		Author:	ParseRule(doc, source.DetailBookAuthorRule, "text"),
		Cover:	FormatURL(ParseRule(doc, source.DetailBookCoverRule, "text"), source.SourceURL),
		Category:	ParseRule(doc, source.DetailBookCategoryRule, "text"),
		Description:	ParseRule(doc, source.DetailBookDescriptionRule, "html"),
	}

	return
}

func (s *detailService) GetDetailNewChapterList(source datamodels.BookSource, doc *html.Node, url string) (newChapterList []datamodels.Chapter) {
	if ncNodes, err := htmlquery.QueryAll(doc, source.DetailNewChapterRule); err == nil {
		for _, node := range ncNodes{
			result := datamodels.Chapter{
				Title: ParseRule(node, source.DetailNewChapterTitleRule, "text"),
				URL:   FormatURL(ParseRule(node, source.DetailNewChapterURLRule, "text"), url),
				Source: source.SourceKey,
			}
			newChapterList = append(newChapterList, result)
		}
	}

	return
}

func (s *detailService) GetDetailChapterList(source datamodels.BookSource, doc *html.Node, url string) (chapterList []datamodels.Chapter) {
	if cNodes, err := htmlquery.QueryAll(doc, source.DetailChapterRule); err == nil {
		for _, node := range cNodes{
			result := datamodels.Chapter{
				Title: ParseRule(node, source.DetailChapterTitleRule, "text"),
				URL:   FormatURL(ParseRule(node, source.DetailChapterURLRule, "text"), url),
				Source: source.SourceKey,
			}
			chapterList = append(chapterList, result)
		}
	}

	return
}

func (s *detailService) GetSourceByKey(key string) (datamodels.BookSource, bool) {
	return s.rep.Select(func(s datamodels.BookSource) bool {
		return s.SourceKey == key
	})
}