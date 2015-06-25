package blog

import "testing"

func TestListFromPage(t *testing.T) {
	anArticle := new(Article)
	articles := []*Article{anArticle, anArticle, anArticle}

	pager := ArticlePager{articles, 2}

	listing, err := pager.ListFromPage(2)
	if err != nil {
		t.Errorf("Error while listing :\n %s", err.Error())
	}
	if len(listing) != 1 {
		t.Errorf("Should have returned 1 article, got %d", len(listing))
	}

	listing, err = pager.ListFromPages(1, 2)
	if err != nil {
		t.Errorf("Error while listing :\n %s", err.Error())
	}
	if len(listing) != 3 {
		t.Errorf("Should have returned 3 article, got %d", len(listing))
	}
}
