package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/")

	if path != "" && path != "index" {
		http.NotFound(w, r)
		return
	}

	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fmt.Println("ecocosts: started")
	assets := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))
	http.Handle("/assets/", assets)
	http.HandleFunc("/", indexHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(fmt.Errorf("Failed to start HTTP server: %v", err))
	}
}
