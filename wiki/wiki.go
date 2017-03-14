package main

import (
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("index.html")).Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.ListenAndServe (":8080", nil);
}
