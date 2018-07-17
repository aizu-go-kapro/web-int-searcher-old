package domain 

import (

)

type (
	Page struct {
		PageID	string
		URL		string
		Text 	string
	}

	PageRepository interface {
		SavePage(db *dbutil.DB, page Page) error
		SavePages(db *dbutil.DB, pages []*Page) error
		Get(db *dbutil.DB, id string) (Page, error)
		Update(db *dbutil.DB, id string) error
		GetDocumentByURL(db *dbutil.DB, url string) (Page, error)
	}
)

func NewPage(pageID string, url string, text string) *Page {
	return &Page{
		PageID: pageID,
		URL: url,
		Text: text,
	}
}