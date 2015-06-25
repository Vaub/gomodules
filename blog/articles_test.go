package blog

import (
	"sort"
	"testing"
	"time"
)

func TestSortByName(t *testing.T) {
	anArticle := Article{"Zeta", "", time.Now()}
	anotherArticle := Article{"Beta", "", time.Now()}

	articles := []*Article{&anArticle, &anotherArticle}

	sort.Sort(ByName(articles))
	if articles[0] != &anotherArticle {
		t.Error("Failed to sort by name")
	}
}
