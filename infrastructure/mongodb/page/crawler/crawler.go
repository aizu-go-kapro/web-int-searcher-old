package crawler

import "github.com/aizu-go-kapro/web-int-searcher/infrastructure/mongodb/page/crawler/crawle"

var pages []*crawle.CrawlePage

func Crawle(url, title string) ([]*crawle.CrawlePage, error) {

	page, err := crawle.NewPage(url, title)
	if err != nil {
		return nil, err
	}
	err = page.GetText()
	if err != nil {
		return nil, err
	}
	err = page.GetLink()
	if err != nil {
		return nil, err
	}
	pages = append(pages, page)
	if len(page.Tolink) == 0 {
		return nil, nil
	}
	for _, pageurl := range page.Tolink {
		_, err := Crawle(pageurl, "")
		if err != nil {
			// return err
		}
	}
	return pages, nil
}
