package document_manegement

import (
	"github.com/aizu-go-kapro/web-int-searcher/crawler/page"
	"github.com/aizu-go-kapro/web-int-searcher/mongo"
)

func SaveDocument(document mongo.Document) error {
	document2, err := mongo.GetDocumentByURL(document.URL)
	if err != nil {
		document.PageID = document2.PageID
		err := mongo.UpdateDocument(document)
		if err != nil {
			return err
		}
	} else {
		err := mongo.SaveDocument(document)
		if err != nil {
			return err
		}
	}

	return nil
}

func SavePages(pages []page.Page) error {
	for _, page := range pages {
		document := mongo.Document{
			page.Id,
			page.Url,
			page.Text,
		}
		err := SaveDocument(document)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDocument(id string) (mongo.Document, error) {
	document, err := mongo.GetDocument(id)
	if err != nil {
		return document, err
	}
	return document, nil
}
