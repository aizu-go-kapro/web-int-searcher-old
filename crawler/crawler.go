package crawler

import (
	"fmt"
	"github.com/aizu-go-kapro/web-int-searcher/crawler/page"
	"log"
)

var pages []page.Page

var sawPages []string

func Crawler() {
	toppages, err := page.LoadTopPage()
	if err != nil {
		log.Println(err)
	}

	for _, page := range toppages {
		crawle(page.Url, page.Title)
	}

	fmt.Println(pages)
}

func crawle(url, title string) error{
	//for _, swapage := range  sawPages {
	//	if swapage == url {
	//		return nil
	//	}
	//}
	page := page.NewPage(url, title)
	err := page.GetText()
	sawPages = append(sawPages, page.Url)
	if err != nil {
		return err
	}
	err = page.GetLink()
	if err != nil {
		return err
	}
	pages = append(pages, page)

	fmt.Println(pages)

	for _, pageurl := range page.Tolink{
		crawle(pageurl, "")
	}



	return nil
}
