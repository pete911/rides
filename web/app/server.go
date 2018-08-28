package app

import (
	"fmt"
	"github.com/husobee/vestigo"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

const layoutPath = "web/template/layout.html"

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
	router.Get("/login", login)
	router.Post("/login", dashboard)

	return router
}

func index(w http.ResponseWriter, _ *http.Request) {
	writeTemplate("index", w, nil)
}

func login(w http.ResponseWriter, _ *http.Request) {
	writeTemplate("login", w, nil)
}

func dashboard(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("login success"))
}

func writeTemplate(templateName string, w http.ResponseWriter, data interface{}) {

	templatePath := fmt.Sprintf("web/template/%s.html", templateName)
	t, err := template.New(path.Base(layoutPath)).ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Printf("%s: parse template files: %v", templateName, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		log.Printf("%s: execute template: %v", templateName, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
