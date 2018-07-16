package mongo

import (
	"github.com/aizu-go-kapro/web-int-searcher/crawler/page"
	mgo "gopkg.in/mgo.v2"
)

func SavePages(pages []page.Page) error {
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("test").C("data")
	for _, page := range pages {
		err := c.Insert(&page)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
