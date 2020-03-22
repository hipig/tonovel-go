package services

import (
	"tonovel/datamodels"
)

type BookService interface {
	GetListByKeyword(keyword string) []*datamodels.BookInfo
	GetInfo(url string, source string) datamodels.BookInfo
	GetChapterList(url string, source string) []datamodels.Chapter
	GetContent(detailURL string, chapterURL string, source string) datamodels.BookContent
}

func NewBookService() BookService {
	return &bookService{}
}

type bookService struct {}

var (
	fService = NewFetcherService()
)

func (s *bookService) GetListByKeyword(keyword string) (itemList []*datamodels.BookInfo) {
	itemList = fService.GetItemList(keyword)
	return
}

func (s *bookService) GetInfo(url string, key string) (info datamodels.BookInfo) {
	info = fService.GetItem(url, key)
	return
}

func (s *bookService) GetChapterList(url string, key string) (chapterList []datamodels.Chapter) {
	chapterList = fService.GetChapterList(url, key)
	return
}

func (s *bookService) GetContent(detailURL string, chapterURL string, key string) (content datamodels.BookContent) {
	content = fService.GetContent(detailURL, chapterURL, key)
	return
}
