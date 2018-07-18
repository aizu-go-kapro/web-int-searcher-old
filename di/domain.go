package di

import (
	"github.com/aizu-go-kapro/web-int-searcher/domain"
	"github.com/aizu-go-kapro/web-int-searcher/infrastructure/mongodb/index"
	"github.com/aizu-go-kapro/web-int-searcher/infrastructure/mongodb/page"
)

var indexRepository domain.IndexRepository

func InjectIndexRepository() domain.IndexRepository {
	if indexRepository != nil {
		return indexRepository
	}

	indexRepository = index.NewRepository()
	return indexRepository
}

var pageRepository domain.PageRepository

func InjectPageRepository() domain.PageRepository {
	if pageRepository != nil {
		return pageRepository
	}

	pageRepository = page.NewRepository()
	return pageRepository
}
