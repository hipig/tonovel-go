package services

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"golang.org/x/net/html"
	"log"
	"strings"
	"sync"
	"tonovel/datamodels"
)

type BookService interface {
	GetListByKeyword(keyword string) []datamodels.SearchItem

	GetDetail(url string, source string) (datamodels.BookInfo, []datamodels.Chapter, []datamodels.Chapter)
	GetDetailInfo(source datamodels.BookSource, doc *html.Node) (info datamodels.BookInfo)
	GetDetailNewChapterList(source datamodels.BookSource, doc *html.Node, url string) (newChapterList []datamodels.Chapter)
	GetDetailChapterList(source datamodels.BookSource, doc *html.Node, url string) (chapterList []datamodels.Chapter)

	GetContent(url string, source string) datamodels.BookContent
}

func NewBookService() BookService {
	return &bookService{}
}

type bookService struct {}

var sourceService  = NewBookSourceService()

func (s *bookService) GetListByKeyword(keyword string) (results []datamodels.SearchItem) {
	var wg sync.WaitGroup
	bookSources :=sourceService.GetAllSource()
	for i := range bookSources{
		wg.Add(1)
		go func(source *datamodels.BookSource) {
			defer wg.Done()
			NewColly(func(resp *colly.Response) {
				doc, err := htmlquery.Parse(strings.NewReader(string(resp.Body)))
				if err != nil {
					log.Fatalln(err)
				}

				nodes := htmlquery.Find(doc, source.SearchItemRule)
				for _, node := range nodes{
					result := datamodels.SearchItem{
						Name:	ParseRule(node, source.SearchItemNameRule, "text"),
						Author:	ParseRule(node, source.SearchItemAuthorRule, "text"),
						Cover:	FormatURL(ParseRule(node, source.SearchItemCoverRule, "text"), source.SourceURL),
						NewChapter:	ParseRule(node, source.SearchItemNewChapterRule, "text"),
						URL:	FormatURL(ParseRule(node, source.SearchItemURLRule, "text"), source.SourceURL),
						Source: source.SourceKey,
					}
					results = append(results, result)
				}
			}).Visit(fmt.Sprintf(source.SearchURL, keyword))
		}(&bookSources[i])
	}
	wg.Wait()
	return
}

func (s *bookService) GetDetail(url string, key string) (info datamodels.BookInfo, newChapterList []datamodels.Chapter, chapterList []datamodels.Chapter) {
	source, ok := sourceService.GetSourceByKey(key)
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

func (s *bookService) GetDetailInfo(source datamodels.BookSource, doc *html.Node) (info datamodels.BookInfo) {
	info = datamodels.BookInfo{
		Name:	ParseRule(doc, source.DetailBookNameRule, "text"),
		Author:	ParseRule(doc, source.DetailBookAuthorRule, "text"),
		Cover:	FormatURL(ParseRule(doc, source.DetailBookCoverRule, "text"), source.SourceURL),
		Category:	ParseRule(doc, source.DetailBookCategoryRule, "text"),
		Description:	ParseRule(doc, source.DetailBookDescriptionRule, "html"),
	}
	return
}

func (s *bookService) GetDetailNewChapterList(source datamodels.BookSource, doc *html.Node, url string) (newChapterList []datamodels.Chapter) {
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

func (s *bookService) GetDetailChapterList(source datamodels.BookSource, doc *html.Node, url string) (chapterList []datamodels.Chapter) {
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

func (s *bookService) GetContent(url string, key string) (content datamodels.BookContent) {
	source, ok := sourceService.GetSourceByKey(key)
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