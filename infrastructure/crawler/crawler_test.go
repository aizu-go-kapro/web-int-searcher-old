package crawler

import (
	"testing"
)

func TestCrawler(t *testing.T) {
	_, err := Crawler()
	if err != nil {
		panic(err)
	}
	// log.Fatal(page)
	_ = t
}
