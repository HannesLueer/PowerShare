package frontend

import (
	"embed"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
)

//go:embed dist/*
var spaFiles embed.FS

// SpaHandler handles the single-page application.
//
// To do this, it tries to find all paths specified in the URL in the file system of the SPA.
// If no matching file is found, the index.html file is delivered. This way it is also possible to link to URLs of the SPA externally or to execute a reload.
func SpaHandler(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("dist", r.URL.Path)

	data, err := spaFiles.ReadFile(filePath)
	if err != nil {
		// serve index.html
		filePath = path.Join("dist", "index.html")
		data, err = spaFiles.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
		}
	}

	var contentType string
	switch filepath.Ext(filePath) {
	case ".html":
		contentType = "text/html"
	case ".js":
		contentType = "text/javascript"
	case ".css":
		contentType = "text/css"
	case ".png":
		contentType = "image/png"
	case ".ico":
		contentType = "image/vnd.microsoft.icon"
	case ".svg":
		contentType = "image/svg+xml"
	default:
		contentType = "text/plain"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Accept-Ranges", "bytes")

	w.Write(data)
}
