package main

import (
	"log"

	"github.com/aizu-go-kapro/web-int-searcher/crawler"
	"github.com/aizu-go-kapro/web-int-searcher/document_manegement"
	"github.com/aizu-go-kapro/web-int-searcher/indexmaker"
)

func main() {
	log.Println("start:\tindex_building_machine")
	log.Println("moving:\tcrawler")
	pages, err := crawler.Crawler()
	if err != nil {
		panic(err)
	}
	log.Println("finished:\tcrawler")
	log.Println("start:\tindex_maker")
	err = indexmaker.MakeIndex(pages)
	if err != nil {
		panic(err)
	}
	log.Println("finished:\tindex_maker")
	log.Println("start:\tdocuments_maker")
	err = document_manegement.SavePages(pages)
	if err != nil {
		panic(err)
	}
	log.Println("finished:\tdocuments_maker")

	log.Println("___________closing___________")
	log.Println("___________  bye  __________")
}
