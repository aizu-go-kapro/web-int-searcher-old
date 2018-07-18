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

func (s *SearchApp) Get(query string) (pages []*mpage.PageCollection, err error) {
	const errtag = "SearchApp.Get() failed"

	var word string
	var pc *mpage.PageCollection
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
		pc = mpage.NewPageCollection(&page)
		pages = append(pages, pc)
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
