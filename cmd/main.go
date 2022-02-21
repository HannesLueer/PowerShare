package main

import (
	"PowerShare/config"
	"PowerShare/frontend"
	"io/fs"
	"log"
	"mime"
	"net/http"
)

func main() {
	// define mime type
	mime.AddExtensionType(".js", "application/javascript")

	// serve files
	mux := http.NewServeMux()
	mux.Handle("/",  http.FileServer(getPWAFileSystem()))

	log.Fatalln(http.ListenAndServeTLS(config.PortStr(), config.GetCertFilePath(), config.GetKeyFilePath(), mux))
}

func getPWAFileSystem() http.FileSystem {
	fsys, err := fs.Sub(frontend.PWA, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}

