package blog

import (
	"errors"

	"github.com/vaub/gomodules/paging"
)

// Blog : a blog using the pager interface
type Blog struct {
	Articles        []Article
	articlesPerPage int
}

// Len : number of articles in the blog
func (blog *Blog) Len() int {
	return len(blog.Articles)
}

// PerPage : number of articles per page
func (blog *Blog) PerPage() int {
	return blog.articlesPerPage
}

// NewBlog : instanciate a new blog with a number of elements per page
func NewBlog(perPage int) (*Blog, error) {
	if perPage <= 0 {
		return nil, errors.New("Cannot create a blog with 0 or less articles per page")
	}

	blog := new(Blog)
	blog.Articles = []Article{}
	blog.articlesPerPage = perPage

	return blog, nil
}

// ListArticlesFromPage : list the articles for a page
func (blog *Blog) ListArticlesFromPage(page int) ([]Article, error) {
	return blog.ListArticlesFromPageRange(page, page)
}

// ListArticlesFromPageRange : list the articles from a page to a page
func (blog *Blog) ListArticlesFromPageRange(from int, to int) ([]Article, error) {
	items, err := paging.ListPageRange(blog, from, to)
	if err != nil {
		return nil, err
	}

	return blog.Articles[items.From:items.To], nil
}
