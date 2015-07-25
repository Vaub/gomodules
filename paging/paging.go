package paging

// Pager : default paging interface
type Pager interface {
	// Len : number of articles
	Len() int
	// Number of article per page
	PerPage() int
}
