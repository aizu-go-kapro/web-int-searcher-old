package crawler

import (
	"fmt"
	"log"

	"github.com/aizu-go-kapro/web-int-searcher/infrastructure/crawler/page"
	uuid "github.com/satori/go.uuid"
)

var pages []page.Page

func Crawler() ([]page.Page, error) {

	toppages, err := page.LoadTopPage()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, page := range toppages {
		crawle(page.Url, page.Title)
	}

	for i := range pages {
		u1, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		pages[i].Id = fmt.Sprint(u1)
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
