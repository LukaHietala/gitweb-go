package main

import (
	"html/template"
	"log"
	"net/http"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) LoadTemplates() {
    t.templates = template.Must(template.ParseGlob("web/views/*.html"))
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
}

func (t *Templates) RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
    t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Templates{}
	t.LoadTemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.RenderTemplate(w, "home.html", nil)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}