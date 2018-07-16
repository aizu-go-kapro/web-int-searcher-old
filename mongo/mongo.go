package mongo

import (
	"github.com/aizu-go-kapro/web-int-searcher/crawler/page"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Index struct {
	Word    string   `json:"word"`
	PageIDs []string `json:"pageIDs"`
}

func SavePages(pages []page.Page) error {
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("test").C("index")
	for _, page := range pages {
		err := c.Insert(&page)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveIndex(index Index) error {
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("test").C("index")
	err = c.Insert(&index)
	if err != nil {
		return err
	}
	return nil
}

func FindPage(word string) ([]string, error) {
	index, err := GetIndex(word)
	if err != nil {
		return nil, err
	}
	return index.PageIDs, nil
}

func UpdateIdex(index Index) error {
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("test").C("index")
	selector := bson.M{"word": index.Word}
	err = c.Update(selector, index)
	if err != nil {
		return err
	}
	return nil
}

func GetIndex(word string) (Index, error) {
	var index Index
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return index, err
	}
	defer session.Close()

	err = session.DB("test").C("index").Find(bson.M{"word": word}).One(&index)
	if err != nil {
		return index, err
	}
	return index, nil
}

type Document struct {
	PageID string `json:"pageID"`
	URL    string `json:"url"`
	Text   string `json:"text"`
}

func SaveDocuments(documents []Document) error {
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("test").C("document")
	for _, document := range documents {
		err := c.Insert(&document)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveDocument(document Document) error {
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("test").C("document")

	err = c.Insert(&document)
	if err != nil {
		return err
	}

	return nil
}

func GetDocument(id string) (Document, error) {
	var document Document
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return document, err
	}
	defer session.Close()

	err = session.DB("test").C("document").Find(bson.M{"pageID": id}).One(&document)
	if err != nil {
		return document, err
	}
	return document, nil
}

func UpdateDocument(document Document) error {
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("test").C("document")
	selector := bson.M{"pageID": document.PageID}
	err = c.Update(selector, document)
	if err != nil {
		return err
	}
	return nil
}

func GetDocumentByURL(url string) (Document, error) {
	var document Document
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		return document, err
	}
	defer session.Close()

	err = session.DB("test").C("document").Find(bson.M{"url": url}).One(&document)
	if err != nil {
		return document, err
	}
	return document, nil
}
