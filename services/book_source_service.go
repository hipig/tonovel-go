package services

import (
	"tonovel/datamodels"
	"tonovel/datasource"
	"tonovel/repositories"
)

type BookSourceService interface {
	GetSourceByKey(key string) (datamodels.BookSource, bool)
	GetAllSource() []datamodels.BookSource
}

func NewBookSourceService() BookSourceService {
	return &bookSourceService{
		repositories.NewBookSourceRepository(datasource.BookSources),
	}
}

type bookSourceService struct {
	rep repositories.BookSourceRepository
}

func (s *bookSourceService) GetSourceByKey(key string) (datamodels.BookSource, bool) {
	return s.rep.Select(func(s datamodels.BookSource) bool {
		return s.SourceKey == key
	})
}

func (s *bookSourceService) GetAllSource() []datamodels.BookSource {
	return s.rep.SelectMany(func(_ datamodels.BookSource) bool {
		return true
	}, -1)
}