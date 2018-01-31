package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	staticFile := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))
	mux.HandleFunc("/", indexHandler)

	sever := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	sever.ListenAndServe()
}

func indexHandler(wr http.ResponseWriter, r *http.Request) {
	templates := []string{
		"static/templates/layout.html",
		"static/templates/index.html",
	}
	temps := template.Must(template.ParseFiles(templates...))
	temps.ExecuteTemplate(wr, "layout", "")
}
