package page

import (
	domain "github.com/aizu-go-kapro/web-int-searcher/domain"
)

type PageCollection struct {
	PageID string `json:"pageID"`
	URL    string `json:"url"`
	Text   string `json:"text"`
}

func NewPageCollection(p *domain.Page) *PageCollection {
	const errtag = "mongodb/page/NewPageCollection failed"

	return &PageCollection{
		PageID: p.PageID,
		URL:    p.URL,
		Text:   p.Text,
	}
}

func (p *PageCollection) Domain() *domain.Page {
	return &domain.Page{
		PageID: p.PageID,
		URL:    p.URL,
		Text:   p.Text,
	}
}
