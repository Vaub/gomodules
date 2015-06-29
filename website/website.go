package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	index string = "index.html"
)

var templates = template.Must(template.ParseFiles(staticResource(index)))

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", defaultHandler)

	http.ListenAndServe(":8080", r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page not found!")
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, index, nil)
}

func staticResource(name string) string {
	return fmt.Sprintf("./static/%s", name)
}
