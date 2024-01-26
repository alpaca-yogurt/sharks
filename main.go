package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

const _index string = "index.html"
const _robots string = "robots.txt"
const _404 string = "404.html"

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t := template.Must(template.ParseFiles(_404))
		t.Execute(w, nil)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	t := template.Must(template.ParseFiles(_index))
	t.Execute(w, nil)
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/robots.txt" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, _robots)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
