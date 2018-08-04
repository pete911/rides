package app

import (
	"fmt"
	"github.com/husobee/vestigo"
	"html/template"
	"log"
	"net/http"
	"time"
)

func NewRidesServer(port int) *http.Server {

	return &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        handler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func handler() http.Handler {

	router := vestigo.NewRouter()

	static := http.FileServer(http.Dir("web/static"))
	router.Handle("/static", http.StripPrefix("/static", static))
	router.Handle("/static/:dir/:file", http.StripPrefix("/static", static))

	router.Get("/", index)

	return router
}

func index(w http.ResponseWriter, _ *http.Request) {

	templatePath := "web/template/index.html"
	t := template.New("index.html")
	t, err := t.ParseFiles(templatePath)
	if err != nil {
		log.Printf("index: parse %s template error: %v", templatePath, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, map[string]string{"title": "rides"}); err != nil {
		log.Printf("index: extecute %s template error: %v", templatePath, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
