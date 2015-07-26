package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaub/gomodules/website/handlers"
)

var (
	templates = template.Must(template.ParseFiles(
		"templates/default.tmpl",
		"templates/index.tmpl"))
)

func main() {
	handlers.InitBlog()

	r := mux.NewRouter().StrictSlash(false)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	r.HandleFunc("/", defaultHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/jquery/").Handler(http.StripPrefix("/jquery/", http.FileServer(http.Dir("./bower_components/jquery/dist/"))))

	api := r.PathPrefix("/api").Subrouter()
	api.Headers("Content-Type", "application/json")

	blogAPI := api.PathPrefix("/blog").Subrouter()
	blogAPI.Methods("GET").HandlerFunc(handlers.BlogHandler)

	http.ListenAndServe(":8080", r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page not found!")
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "default.tmpl", nil)
}

func staticResource(name string) string {
	return fmt.Sprintf("./static/%s", name)
}
