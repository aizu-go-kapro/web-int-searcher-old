package main

import (
	"net/http"
)

func main() {
	

	bmApp := application.NewBuildingMachine(
		di.InjectDBnjectDB(),
		di.InjectIndexRepository(),
		di.InjectPageRepository()
	)
	
	go bmApp.Run()

	http.HandlFunc("/search", SearchHandler)
	http.ListenAndServe(":3000", nil)
}


func SearchHandler(w http.ResponseWriter, r *http.Request) {
	res := appllication.SearchApp(
		//repositoryとかDBのインスタンスをDIする
		r,
	)
	fmt.Fprintf(w, res)
}