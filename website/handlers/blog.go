package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vaub/gomodules/blog"
)

const (
	defaultPerPage int = 5
)

// BlogEntries : JSON object
type BlogEntries struct {
	Page          int      `json:"page"`
	NumberOfPages int      `json:"number_of_pages"`
	Articles      []string `json:"articles"`

	Path string `json:"path"`
}

// InitBlog : initialize a new blog from a path
// Will default to empty if it cannot find or fetch from the path
func InitBlog(path string, perPage int) *blog.Blog {
	var err error
	var blogPager *blog.Blog

	blogPager, err = blog.FetchFromPath(path)
	if err != nil {
		blogPager, _ = blog.NewBlog(perPage)
	}

	fmt.Printf("Created blog with %d articles", len(blogPager.Articles))
	return blogPager
}

// BlogHandler : a common blog handler function
func BlogHandler(blogPager *blog.Blog, path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := 1
		articles, err := blogPager.ListArticlesFromPage(page)
		if err != nil {
			articles = []blog.Article{}
		}

		articlesInfo := []string{}
		for _, article := range articles {
			articlesInfo = append(articlesInfo, article.Path)
		}

		entries := BlogEntries{page, page, articlesInfo, path}
		json.NewEncoder(w).Encode(entries)
	}
}
