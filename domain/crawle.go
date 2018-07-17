package domain

import (
	"github.com/aizu-go-kapro/web-int-searcher/infrastructure/crawler"
	"github.com/aizu-go-kapro/web-int-searcher/infrastructure/crawler/page"
)

func GetRawPages() ([]page.Page, error) {
	pages, err := crawler.Crawler()
	return page, err
}
