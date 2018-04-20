package crawler

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Link struct {
	Url   string
	Title string
	Date  int64
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

func SearchDataFromURL(link Link) (Link, error) {
	doc, err := goquery.NewDocument(link.Url)
	if err != nil {
		fmt.Println(err, doc)
		return link, err
	}
	link.Title = doc.Find("title").Text()
	link.Date = time.Now().Unix()
	return link, nil
}
