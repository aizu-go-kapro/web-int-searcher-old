package main

import (
	"encoding/json"
	"log"
	"net/http"

	application "github.com/aizu-go-kapro/web-int-searcher/application"
	"github.com/aizu-go-kapro/web-int-searcher/di"
)

func main() {
	/*
		bmApp := application.BuildingMachine(
			// diして
		)

		go bmApp
	*/

	http.HandleFunc("/search", SearchHandler)
	http.ListenAndServe(":3000", nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	searchApp := application.NewSearchApp(
		//repositoryとかDBのインスタンスをDIする
		di.InjectIndexRepository(),
		di.InjectPageRepository(),
		di.InjectDB(),
	)
	// セキュリティがやばい
	q := r.URL.Query()
	pages, err := searchApp.Get(q)
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(pages)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(res)
	if err != nil {
		panic(err)
	}
}
