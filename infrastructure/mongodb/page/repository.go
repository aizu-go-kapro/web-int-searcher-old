package page

import (
	"fmt"
	"log"

	"github.com/aizu-go-kapro/web-int-searcher/domain"
	"github.com/aizu-go-kapro/web-int-searcher/infrastructure/mongodb/page/crawler"
	"github.com/aizu-go-kapro/web-int-searcher/infrastructure/mongodb/page/crawler/crawle"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

func (r *Repository) NewRepository() *Repository {
	return &Repository{}
}

// TODO: domain/page.go PageInterfaceを実装する。

func (r *Repository) SavePage(session mgo.Session, p domain.Page) error {
	c := session.DB("test").C("page")
	err := c.Insert(&p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SavePages(session mgo.Session, pages []domain.Page) error {
	c := session.DB("test").C("page")
	for _, page := range page {
		err := r.SavePage(ms, page)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) Get(session mgo.Session, id string) (domain.Page, error) {
	var page domain.Page
	err := session.DB("test").C("page").Find(bson.M{"pageID": id}).One(&page)
	if err != nil {
		return page, err
	}

	return page, nil
}

func (r *Repository) Update(session mgo.Session, p domain.Page) error {
	c := session.DB("test").C("page")
	selector := bson.M{"pageID": p.PageID}
	err = c.Update(selector, document)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetDocumentByURL(session mgo.Session, url string) (domain.Page, error) {
	var page domain.Page
	err := session.DB("test").C("page").Find(bson.M{"url": url}).One(&page)
	if err != nil {
		return page, err
	}
	return page, nil
}

func (r *Repository) GetFromCrawler() ([]*domain.Page, error) {
	toppages, err := crawle.LoadTopPage()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	pages := []*crawle.CrawlePage{}
	for _, page := range toppages {
		p, err := crawler.Crawle(page.Url, page.Title)
		if err != nil {
			return nil, err
		}
		pages = append(pages, p...)
	}

	for i := range pages {
		u1, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		pages[i].Id = fmt.Sprint(u1)
	}

	result_page := []*domain.Page{}
	for _, page := range pages {
		result_page = append(result_page, domain.NewPage(
			page.Id,
			page.Url,
			page.Text,
			page.Title,
		))
	}

	return result_page, nil
}
