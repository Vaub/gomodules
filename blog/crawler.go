package blog

import (
	"errors"
	"io/ioutil"
	"path"
)

const (
	defaultPerPage = 5
)

// FetchFromPath : create and fill the blog from a path
func FetchFromPath(dirPath string) (*Blog, error) {
	blog, err := NewBlog(defaultPerPage)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, errors.New("Could not read from path.")
	}

	for _, file := range content {
		fileName := file.Name()

		if isFileHTML(fileName) {
			article := Article{path.Base(fileName), fileName, file.ModTime()}
			blog.Articles = append(blog.Articles, article)
		}
	}

	return blog, nil
}

func isFileHTML(filePath string) bool {
	return path.Ext(filePath) == ".html"
}
