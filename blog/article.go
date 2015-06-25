package blog

import "time"

// Article : represents an article (name, path, creation)
type Article struct {
	name     string
	path     string
	modified time.Time
}

// ByName : sort articles by name
type ByName []*Article

func (s ByName) Len() int           { return len(s) }
func (s ByName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByName) Less(i, j int) bool { return s[i].name < s[j].name }

// ByModified : sort articles by date modified
type ByModified []*Article

func (s ByModified) Len() int           { return len(s) }
func (s ByModified) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByModified) Less(i, j int) bool { return s[i].modified.Before(s[j].modified) }
