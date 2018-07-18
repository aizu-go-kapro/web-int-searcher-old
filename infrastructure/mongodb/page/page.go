package page

import (
	"github.com/aizu-go-kapro/web-int-searcher/domain"
)

type PageCollection struct {
	PageID string `json:"pageID"`
	URL    string `json:"url"`
	Text   string `json:"text"`
}

func NewPageColloction(p *domain.Page) (*PageCollection, error) {
	const errtag = "mongodb/page/NewPageCollection failed"

	return &PageCollection{
		PageID: p.PageID,
		URL:    p.URL,
		Text:   p.Text,
	}, nil
}

func (p *PageCollection) Domain() *domain.Page {
	return &domain.Page{
		PageID: p.PageID,
		URL:    p.URL,
		Text:   p.Text,
	}
}
