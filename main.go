package main

import (
	"log"
	"net/http"
	"html/template"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	http.HandleFunc("/foo", fooHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {

	t := template.New("index.html")
	t, err := t.ParseFiles("web/template/index.html")
	if err != nil {
		// TODO
		log.Printf("error: %v", err)
		return
	}

	if err := t.Execute(w, map[string]string{"title": "rides"}); err != nil {
		// TODO
		log.Printf("error: %v", err)
	}
}
