package services

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"sync"
	"tonovel/datamodels"
)

type BookService interface {
	GetListByKeyword(keyword string) []datamodels.BookInfo
	getItemSearch(source *datamodels.BookSource, doc *html.Node) datamodels.BookInfo
	GetInfo(url string, source string) datamodels.BookInfo
	getItemInfo(source *datamodels.BookSource, doc *html.Node) datamodels.BookInfo
	GetChapterList(url string, source string) []datamodels.Chapter
	getChapterList(source *datamodels.BookSource, doc *html.Node, url string) []datamodels.Chapter
	GetContent(detailURL string, chapterURL string, source string) datamodels.BookContent
	getContent(source *datamodels.BookSource, doc *html.Node, url string) datamodels.BookContent
}

func NewBookService() BookService {
	return &bookService{}
}

type bookService struct {}

var sourceService  = NewBookSourceService()

func (s *bookService) GetListByKeyword(keyword string) (itemList []datamodels.BookInfo) {
	var wg sync.WaitGroup
	bookSources :=sourceService.GetAllSource()
	for i := range bookSources{
		wg.Add(1)
		go func(source *datamodels.BookSource) {
			defer wg.Done()
			NewColly(func(doc *html.Node) {
				if nodes, err := htmlquery.QueryAll(doc, source.SearchItemRule); err == nil {
					for _, node := range nodes {
						itemList = append(itemList, s.getItemSearch(source, node))
					}
				}
			}).Visit(fmt.Sprintf(source.SearchURL, keyword))
		}(&bookSources[i])
	}
	wg.Wait()
	return
}

func (s *bookService) getItemSearch(source *datamodels.BookSource, doc *html.Node) (item datamodels.BookInfo) {
	item = datamodels.BookInfo{
		Name:	ParseRule(doc, source.SearchItemNameRule, "text"),
		Author:	ParseRule(doc, source.SearchItemAuthorRule, "text"),
		Cover:	FormatURL(ParseRule(doc, source.SearchItemCoverRule, "text"), source.SourceURL),
		NewChapter:	ParseRule(doc, source.SearchItemNewChapterRule, "text"),
		URL:	FormatURL(ParseRule(doc, source.SearchItemURLRule, "text"), source.SourceURL),
		Source: source.SourceKey,
	}
	return
}

func (s *bookService) GetInfo(url string, key string) (info datamodels.BookInfo) {
	source, ok := sourceService.GetSourceByKey(key)
	if !ok {
		info = datamodels.BookInfo{}
		return
	}

	NewColly(func(doc *html.Node) {
		info = s.getItemInfo(&source, doc)
	}).Visit(url)
	return
}

func (s *bookService) getItemInfo(source *datamodels.BookSource, doc *html.Node) (info datamodels.BookInfo) {
	info = datamodels.BookInfo{
		Name:	ParseRule(doc, source.DetailBookNameRule, "text"),
		Author:	ParseRule(doc, source.DetailBookAuthorRule, "text"),
		Cover:	FormatURL(ParseRule(doc, source.DetailBookCoverRule, "text"), source.SourceURL),
		Category:	ParseRule(doc, source.DetailBookCategoryRule, "text"),
		Description:	ParseRule(doc, source.DetailBookDescriptionRule, "html"),
	}
	return
}

func (s *bookService) GetChapterList(url string, key string) (chapterList []datamodels.Chapter) {
	source, ok := sourceService.GetSourceByKey(key)
	if !ok {
		chapterList = []datamodels.Chapter{}
		return
	}

	var chapterListURL string
	NewColly(func(doc *html.Node) {
		if source.DetailChapterListURLRule != "" {
			chapterListURL = ParseRule(doc, source.DetailChapterListURLRule, "text")
		}else {
			chapterList = s.getChapterList(&source, doc, url)
			return
		}
	}).Visit(url)

	if chapterListURL != "" {
		NewColly(func(doc *html.Node) {
			chapterList = s.getChapterList(&source, doc, chapterListURL)
		}).Visit(chapterListURL)
	}
	return
}

func (s *bookService) getChapterList(source *datamodels.BookSource, doc *html.Node, url string) (chapterList []datamodels.Chapter) {
	if cNodes, err := htmlquery.QueryAll(doc, source.DetailChapterRule); err == nil {
		for _, node := range cNodes{
			result := datamodels.Chapter{
				Title: ParseRule(node, source.DetailChapterTitleRule, "text"),
				ChapterURL:   FormatURL(ParseRule(node, source.DetailChapterURLRule, "text"), url),
				Source: source.SourceKey,
			}
			chapterList = append(chapterList, result)
		}
	}
	return
}

func (s *bookService) GetContent(detailURL string, chapterURL string, key string) (content datamodels.BookContent) {
	source, ok := sourceService.GetSourceByKey(key)
	if !ok {
		content = datamodels.BookContent{}
		return
	}

	NewColly(func(doc *html.Node) {
		content = s.getContent(&source, doc, detailURL)
	}).Visit(chapterURL)
	return
}

func (s *bookService) getContent(source *datamodels.BookSource, doc *html.Node, url string) (content datamodels.BookContent) {
	content = datamodels.BookContent{
		Title:	ParseRule(doc, source.ContentTitleRule, "text"),
		Text:	ParseRule(doc, source.ContentTextRule, "html"),
		DetailURL: url,
		PreviousURL:	FormatURL(ParseRule(doc, source.ContentPreviousURLRule, "text"), source.SourceURL),
		NextURL:	FormatURL(ParseRule(doc, source.ContentNextURLRule, "text"), source.SourceURL),
		Source: source.SourceKey,
	}
	return content
}