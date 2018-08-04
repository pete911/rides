package main

import (
	"github.com/pete911/rides/web/app"
	"log"
)

func main() {

	//router.Handle("/swagger", http.StripPrefix("/swagger", swaggerUi))
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	//http.HandleFunc("/foo", fooHandler)
	server := app.NewRidesServer(8080)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

//func fooHandler(w http.ResponseWriter, r *http.Request) {
//
//	t := template.New("index.html")
//	t, err := t.ParseFiles("web/template/index.html")
//	if err != nil {
//		// TODO
//		log.Printf("error: %v", err)
//		return
//	}
//
//	if err := t.Execute(w, map[string]string{"title": "rides"}); err != nil {
//		// TODO
//		log.Printf("error: %v", err)
//	}
//}
