package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"gitweb-go/config"
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

	// if config file exists, override default one
	var configPath string
	flag.StringVar(&configPath, "config", "config.yml", "Path to configuration file")
	flag.Parse()

	config, err := config.LoadConfig(configPath)

	if err != nil {
		log.Printf("Error loading config: %s\n", err)
	}

	// if environment variables are set, override config and default values :)

	if reposPath := os.Getenv("REPOS_PATH"); reposPath != "" {
		config.ReposPath = reposPath
		log.Printf("Overriding repos path with env variable: %s\n", config.ReposPath)
	}
	if port := os.Getenv("PORT"); port != "" {
		config.Port = port
		log.Printf("Overriding port with env variable: %s\n", config.Port)
	}

	log.Printf("Starting server on :%s with repos path %s\n", config.Port, config.ReposPath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.RenderTemplate(w, "home.html", nil)
	})

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}