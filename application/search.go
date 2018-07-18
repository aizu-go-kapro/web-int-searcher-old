package application

import (
	"github.com/aizu-go-kapro/web-int-searcher/domain"
	mpage "github.com/aizu-go-kapro/web-int-searcher/infrastructure/mongodb/page"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

type SearchApp struct {
	IndexRepo domain.IndexRepository
	PageRepo  domain.PageRepository
	Session   *mgo.Session
}

func NewSearchApp(
	index domain.IndexRepository,
	page domain.PageRepository,
	session *mgo.Session,
) *SearchApp {
	return &SearchApp{
		IndexRepo: index,
		PageRepo:  page,
		Session:   session,
	}
}

func (s *SearchApp) Get(query string) ([][]*mpage.PageCollection, error) {
	const errtag = "SearchApp.Get() failed"

	words, err := parseQuery(query)
	if err != nil {
		return nil, errors.Wrap(err, errtag)
	}

	indexs := make([]*domain.Index, len(words))
	for _, word := range words {
		index, err := s.IndexRepo.Get(s.Session, word)
		if err != nil {
			return nil, errors.Wrap(err, errtag)
		}
		indexs = append(indexs, index)
	}
	if err != nil {
		return nil, errors.Wrap(err, errtag)
	}
	pages := make([][]*mpage.PageCollection, len(indexs))
	for i, index := range indexs {
		for _, id := range index.PageIDs {
			page, err := s.PageRepo.Get(s.Session, id)
			if err != nil {
				return nil, errors.Wrap(err, errtag)
			}
			pc := mpage.NewPageCollection(&page)
			pages[i] = append(pages[i], pc)
		}
		if err != nil {
			return nil, errors.Wrap(err, errtag)
		}
	}

	return pages, nil
}

//
func parseQuery(query string) ([]string, error) {

	return nil, nil
}
