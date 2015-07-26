package blog

import (
	"fmt"
	"os"
	"testing"
)

const (
	aFile       string = "a_file.html"
	anotherFile string = "another_file.html"

	dir          string = "testing"
	imaginaryDir string = "thisisnotsupposedtoexists"
)

func TestFetchFromPath(t *testing.T) {
	prepareFiles()

	blog, err := FetchFromPath(dir)
	if err != nil {
		t.Errorf("Error while creating the pager :\n%s", err.Error())
	}

	if len(blog.Articles) != 2 {
		t.Errorf("Found %d articles, expected 2", len(blog.Articles))
	}

	deleteFiles()
}

func TestFetchFromInvalidPath(t *testing.T) {
	_, err := FetchFromPath(imaginaryDir)

	if err == nil {
		t.Errorf("Expected an error, does %s exists?", imaginaryDir)
	}
}

func prepareFiles() {
	os.Mkdir(dir, os.ModePerm)

	os.Create(file(aFile))
	os.Create(file(anotherFile))
}

func deleteFiles() {
	os.Remove(file(aFile))
	os.Remove(file(anotherFile))

	os.Remove(dir)
}

func file(name string) string {
	return fmt.Sprintf("%s/%s", dir, name)
}
