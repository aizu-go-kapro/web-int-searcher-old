package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aizu-go-kapro/web-int-searcher/di"
)

func main() {
	/*
		bmApp := application.BuildingMachine(
			// diして
		)

		go bmApp
	*/

	http.HandlFunc("/search", SearchHandler)
	http.ListenAndServe(":3000", nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	searchApp := appllication.NewSearchApp(
		//repositoryとかDBのインスタンスをDIする
		di.InjectIndexRepository(),
		di.InjectPageRepository(),
		di.InjectDB(),
	)

	page, err := searchApp.Get(q)
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(page)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, res)
}
