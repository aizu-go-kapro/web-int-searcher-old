package domain

import (
	"gopkg.in/mgo.v2"
)

type (
	Page struct {
		PageID string
		URL    string
		Title  string
		Text   string
	}

	PageRepository interface {
		SavePage(ms *mgo.Session, page Page) error
		SavePages(ms *mgo.Session, pages []*Page) error
		Get(ms *mgo.Session, id string) (Page, error)
		Update(ms *mgo.Session, p Page) error
		GetDocumentByURL(ms *mgo.Session, url string) (Page, error)
		GetFromCrawler() ([]*Page, error)
	}
)

func NewPage(pageID string, url string, text string, title string) *Page {
	return &Page{
		PageID: pageID,
		URL:    url,
		Text:   text,
		Title:  title,
	}
}
