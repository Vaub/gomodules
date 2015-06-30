package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaub/gomodules/blog"
	"github.com/vaub/gomodules/website/handlers"
)

const (
	index string = "index.html"
)

var (
	blogPager = blog.NewArticlePager()
	templates = template.Must(template.ParseFiles(staticResource(index)))
)

func main() {
	blogPager, _ = blog.FetchFromPath("static/articles")
	handlers.InitBlog()

	r := mux.NewRouter().StrictSlash(false)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	r.HandleFunc("/", defaultHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

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
	templates.ExecuteTemplate(w, index, nil)
}

func staticResource(name string) string {
	return fmt.Sprintf("./static/%s", name)
}
