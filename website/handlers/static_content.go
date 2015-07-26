package handlers

import "net/http"

// StaticContent : handler to server static/file content trough the website
// useful for CSS/JS/HTML
func StaticContent(prefix string, path string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.Dir(path)))
}
