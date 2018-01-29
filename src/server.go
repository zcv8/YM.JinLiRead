package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	staticFile := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))
	mux.HandleFunc("/", loginHandler)

	sever := &http.Server{
		Addr:    "0.0.0.0:8011",
		Handler: mux,
	}
	sever.ListenAndServe()
}

func loginHandler(wr http.ResponseWriter, r *http.Request) {
	templates := []string{
		"static/templates/login.html",
	}
	temps := template.Must(template.ParseFiles(templates...))
	temps.ExecuteTemplate(wr, "login", "")
}
