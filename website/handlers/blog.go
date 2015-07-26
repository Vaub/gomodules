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
}

var blogPager *blog.Blog

func InitBlog() {
	var err error

	blogPager, err = blog.FetchFromPath("static/articles")
	if err != nil {
		blogPager, _ = blog.NewBlog(defaultPerPage)
	}

	fmt.Printf("Created blog with %d articles", len(blogPager.Articles))
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	page := 1
	articles, err := blogPager.ListArticlesFromPage(page)
	if err != nil {
		articles = []blog.Article{}
	}

	articlesInfo := []string{}
	for _, article := range articles {
		articlesInfo = append(articlesInfo, article.Path)
	}

	entries := BlogEntries{page, page, articlesInfo}
	json.NewEncoder(w).Encode(entries)
}
