package page

import (
	"github.com/aizu-go-kapro/web-int-searcher/domain"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

func (r *Repository) NewRepository() {
	return &Repository{}
}

// TODO: domain/page.go PageInterfaceを実装する。

func (r *Repository) SavePage(ms mongoutil.Session, p domain.Page) error {
	c := ms.C("page")
	err := c.Insert(&p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SavePages(ms mongoutil.Session, pages []domain.Page) error {
	c := ms.C("page")
	for _, page := range page {
		err := r.SavePage(ms, page)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) Get(ms mongoutil.Session, id string) (domain.Page, error) {
	var page domain.Page{}
	err := ms.C("page").Find(bson.M{"pageID": id}).One(&page)
	if err != nil {
		return page, err
	}

	return page, nil
}


func (r *Repository) Update(ms mongoutil.Session, p domain.Page) error {
	c := ms.C("page")
	selector := bson.M{"pageID": p.PageID}
	err = c.Update(selector, document)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetDocumentByURL(ms mongoutil.Session, url string) (domain.Page, error) {
	var page domain.Page{}
	err = ms.C("document").Find(bson.M{"url": url}).One(&page)
	if err != nil {
		return document, err
	}
	return nil
}
