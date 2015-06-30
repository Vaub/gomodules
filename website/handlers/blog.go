package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vaub/gomodules/blog"
)

// BlogEntries : JSON object
type BlogEntries struct {
	Page          int      `json:"page"`
	NumberOfPages int      `json:"number_of_pages"`
	Articles      []string `json:"articles"`
}

var pager *blog.ArticlePager

func InitBlog() {
	var err error

	pager, err = blog.FetchFromPath("static/articles")
	if err != nil {
		pager = blog.NewArticlePager()
	}

	pager.SortArticles(blog.SortByModified)
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	page := 1
	articles, err := pager.ListFromPage(page)
	if err != nil {
		articles = []*blog.Article{}
	}

	articlesInfo := []string{}
	for _, article := range articles {
		articlesInfo = append(articlesInfo, article.Path)
	}

	entries := BlogEntries{page, page, articlesInfo}
	json.NewEncoder(w).Encode(entries)
}
