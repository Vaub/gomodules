package paging

import "errors"

// Range : represents a pager's page range
type Range struct {
	// From : index to start from
	From int
	// To : index+1 to end to
	To int
}

// ListPage : list all articles within a page
func ListPage(pager Pager, page int) (Range, error) {
	return ListPageRange(pager, page, page)
}

// ListPageRange : list the articles from page start to page end
// if the end page doesn't exist, it will range to the last element
func ListPageRange(pager Pager, start int, end int) (Range, error) {
	startAt := (start - 1) * pager.PerPage()
	endAt := ((end - 1) * pager.PerPage()) + pager.PerPage()

	if pager.PerPage() <= 0 {
		return Range{}, errors.New("Cannot have 0 or less item per page")
	}

	if start <= 0 || startAt > (pager.Len()-1) || end < start {
		return Range{}, errors.New("Invalid page range")
	}

	if endAt > pager.Len() {
		endAt = pager.Len()
	}

	return Range{startAt, endAt}, nil
}
