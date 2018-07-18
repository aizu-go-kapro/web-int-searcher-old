package main

import (
	"fmt"
	"net/http"

	"github.com/aizu-go-kapro/web-int-searcher/di"
)

func main() {
	// bmApp := application.NewBuildingMachine(
	// 	di.InjectDB(),
	// 	di.InjectIndexRepository(),
	// 	di.InjectPageRepository(),
	// )

	// go bmApp.Run()

	http.HandleFunc("/search", SearchHandler)
	http.ListenAndServe(":3000", nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	searchApp := appllication.NewSearchApp(
		//repositoryとかDBのインスタンスをDIする
		di.InjectIndexRepository(),
		di.InjectPageRepository(),
		di.InjectDB(),
	)
	//TODO: 受け取ったクエリを文字列に
	page := searchApp.Get(q)
	// TODO: pageをjsonのレスポンスにする
	fmt.Fprintf(w, res)
}
