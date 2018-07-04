package crawler

import (
	"fmt"
	"github.com/aizu-go-kapro/web-int-searcher/crawler/page"
	"log"
)

func Crawler() {
	toppages, err := page.LoadTopPage()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(toppages)
}
