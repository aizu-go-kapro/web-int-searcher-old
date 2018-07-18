package domain

import mgo "gopkg.in/mgo.v2"

type (
	Page struct {
		PageID string
		URL    string
		Title  string
		Text   string
	}

	PageRepository interface {
		SavePage(session *mgo.Session, page Page) error
		SavePages(session *mgo.Session, pages []*Page) error
		Get(session *mgo.Session, id string) (Page, error)
		Update(session *mgo.Session, page Page) error
		GetDocumentByURL(session *mgo.Session, url string) (Page, error)
		GetFromCrawler() ([]Page, error)
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
