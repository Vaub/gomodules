package blog

import (
	"errors"
	"fmt"
	"sort"
)

// ArticlePager : a pager containing articles
type ArticlePager struct {
	articles []*Article
	perPage  int
}

const (
	// SortByName : will sort the slice by article name
	SortByName = iota
	// SortByModified : will sort the slice by article modified time
	SortByModified
)

const (
	// DefaultNumberOfArticlesPerPage : default number of articles per page
	DefaultNumberOfArticlesPerPage = 5
)

// NewArticlePager : creates a default pager
func NewArticlePager() *ArticlePager {
	pager := new(ArticlePager)
	pager.articles = []*Article{}
	pager.perPage = DefaultNumberOfArticlesPerPage

	return pager
}

// ListFromPage : list the articles at the specified page
func (pager *ArticlePager) ListFromPage(page int) ([]*Article, error) {
	return pager.ListFromPages(page, 1)
}

// ListFromPages : List articles starting at a page # for # of pages
func (pager *ArticlePager) ListFromPages(page int, pages int) ([]*Article, error) {
	var articles []*Article

	if pager.perPage <= 0 {
		return articles, errors.New("Cannot have 0 or less item per pages")
	}

	var numberOfPages = pager.numberOfPages()

	if page <= 0 || (page+(pages-1)) > numberOfPages {
		return articles, fmt.Errorf("Pages %d to %d does not exists", page, page+(pages-1))
	}

	startAt := (page - 1) * pager.perPage
	endAt := startAt

	maxNbOfArticles := pager.perPage * pages

	if (pager.numberOfArticles() - 1) < startAt+maxNbOfArticles {
		endAt = pager.numberOfArticles()
	} else {
		endAt += maxNbOfArticles
	}

	if startAt == endAt {
		return append(articles, pager.articles[startAt]), nil
	}

	return pager.articles[startAt:endAt], nil
}

// SortArticles : sort articles from predefined presets
func (pager *ArticlePager) SortArticles(sortBy int) {
	switch sortBy {
	case SortByName:
		sort.Sort(ByName(pager.articles))
		return
	case SortByModified:
		sort.Sort(ByModified(pager.articles))
		return
	}
}

func (pager *ArticlePager) numberOfPages() int {
	nbOfArticles := pager.numberOfArticles()
	nbOfPages := nbOfArticles / pager.perPage

	if nbOfArticles%pager.perPage == 0 {
		return nbOfPages
	}

	return nbOfPages + 1
}

func (pager *ArticlePager) numberOfArticles() int {
	return len(pager.articles)
}
