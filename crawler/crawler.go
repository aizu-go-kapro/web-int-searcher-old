package crawler

import (
	"log"

	"github.com/aizu-go-kapro/web-int-searcher/crawler/page"
)

var pages []page.Page

func Crawler() ([]page.Page, error) {
	toppages, err := page.LoadTopPage()
	if err != nil {
		log.Println(err)
	}
	// TODO
	// PageID を登録
	for _, page := range toppages {
		crawle(page.Url, page.Title)
	}
	return pages, nil
}

func crawle(url, title string) error {
	page := page.NewPage(url, title)
	err := page.GetText()
	if err != nil {
		return err
	}
	err = page.GetLink()
	if err != nil {
		return err
	}
	pages = append(pages, page)
	if len(page.Tolink) == 0 {
		return nil
	}
	for _, pageurl := range page.Tolink {
		crawle(pageurl, "")

	}

	return nil
}
