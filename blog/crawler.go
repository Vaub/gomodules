package blog

import (
	"errors"
	"io/ioutil"
	"path"
)

// FetchFromPath : create and fill the pager from a path
func FetchFromPath(dirPath string) (*ArticlePager, error) {
	pager := NewArticlePager()

	content, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, errors.New("Could not read from path.")
	}

	for _, file := range content {
		fileName := file.Name()

		if isFileHTML(fileName) {
			article := &Article{path.Base(fileName), fileName, file.ModTime()}
			pager.articles = append(pager.articles, article)
		}
	}

	return pager, nil
}

func isFileHTML(filePath string) bool {
	return path.Ext(filePath) == ".html"
}
