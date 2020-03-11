package services

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
	"sync"
	"tonovel/datamodels"
	"tonovel/datasource"
	"tonovel/repositories"
)

type SearchService interface {
	SearchByName(name string) []datamodels.SearchItem
	GetAllSource() []datamodels.BookSource
}

func NewSearchService() SearchService {
	return &searchService{
		repositories.NewBookSourceRepository(datasource.BookSources),
	}
}

type searchService struct {
	rep repositories.BookSourceRepository
}

func (s *searchService) SearchByName(name string) (results []datamodels.SearchItem) {
	var wg sync.WaitGroup
	bookSources :=s.GetAllSource()
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
			}).Visit(fmt.Sprintf(source.SearchURL, name))
		}(&bookSources[i])
	}
	wg.Wait()
	return
}

func (s *searchService) GetAllSource() []datamodels.BookSource {
	return s.rep.SelectMany(func(_ datamodels.BookSource) bool {
		return true
	}, -1)
}