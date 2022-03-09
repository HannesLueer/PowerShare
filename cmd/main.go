package main

import (
	"PowerShare/config"
	"PowerShare/frontend"
	"PowerShare/handler/locations"
	"github.com/gorilla/mux"
	"io/fs"
	"log"
	"mime"
	"net/http"
)

func main() {
	// define mime type
	mime.AddExtensionType(".js", "application/javascript")

	r := mux.NewRouter()

	// serve frontend
	frontend := r.PathPrefix("/").Subrouter()
	frontend.Handle("/",  http.FileServer(getPWAFileSystem()))

	// serve api
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/locations", locations.LocationsHandler)

	log.Fatalln(http.ListenAndServeTLS(config.PortStr(), config.GetCertFilePath(), config.GetKeyFilePath(), r))
}

func getPWAFileSystem() http.FileSystem {
	fsys, err := fs.Sub(frontend.PWA, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}

