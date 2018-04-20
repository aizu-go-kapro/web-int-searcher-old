package crawler

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Link struct {
	Url string
}

func GetALLURL(url string) ([]Link, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err, doc)
		return nil, err
	}
	links := []Link{}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		links = append(
			links, Link{
				Url: link,
			},
		)
	})
	return links, nil
}
