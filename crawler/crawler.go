package crawler

import (
	"log"

	"github.com/aizu-go-kapro/web-int-searcher/crawler/page"
	mgo "gopkg.in/mgo.v2"
)

var pages []page.Page

func Crawler() {
	toppages, err := page.LoadTopPage()
	if err != nil {
		log.Println(err)
	}

	for _, page := range toppages {
		crawle(page.Url, page.Title)
	}

	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("data")
	for _, page := range pages {
		err := c.Insert(&page)
		if err != nil {
			panic(err)
		}
	}
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
