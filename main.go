package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/home.html"))
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}
