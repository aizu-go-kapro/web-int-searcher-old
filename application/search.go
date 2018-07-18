package application

import (
	"github.com/aizu-go-kapro/web-int-searcher/domain"
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

func (s *SearchApp) Get(query string) (pages []*domain.Page, err error) {
	const errtag = "SearchApp.Get() failed"

	var word string
	var page domain.Page
	var index *domain.Index

	word, err = parseQuery(query)
	if err != nil {
		return nil, errors.Wrap(err, errtag)
	}

	index, err = s.IndexRepo.Get(s.Session, word)
	if err != nil {
		return nil, errors.Wrap(err, errtag)
	}

	for _, id := range index.PageIDs {
		page, err = s.PageRepo.Get(s.Session, id)
		pages = append(pages, &page)
	}
	if err != nil {
		return nil, errors.Wrap(err, errtag)
	}

	return pages, nil
}

//
func parseQuery(query string) (string, error) {
	return "", nil
}
