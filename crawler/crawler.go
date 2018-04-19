package crawler

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func GetALLURL(doc *goquery.Document) error {
	urls := doc.Find("a").Text()
	fmt.Println(urls)
	return nil
}

func GetDoc(url string) (*goquery.Document, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
