package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := ` 
	<span class="label success" id="result">Success!</span>`

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, htmlContent)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api", apiHandler)

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
