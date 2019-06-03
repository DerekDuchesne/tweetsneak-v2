package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var (
	indexTpl = template.Must(template.ParseFiles("index.html"))
)

func main() {
	http.HandleFunc("/", index)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := indexTpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		panic(err)
	}
}
