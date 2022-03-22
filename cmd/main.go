package main

import (
	"PowerShare/config"
	"PowerShare/database"
	"PowerShare/frontend"
	"PowerShare/handler/chargingStation"
	"PowerShare/helper"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// define mime type
	mime.AddExtensionType(".js", "application/javascript")

	// create default directories and files
	createDefaultDirs()

	// load .env files
	err := godotenv.Load(
		filepath.Join(config.GetConfigFilePath(), "db.env"),
		filepath.Join(config.GetConfigFilePath(), "server.env"),
	)
	if err != nil {
		log.Fatal(err)
	}

	var DBHost 	 	= os.Getenv("DB_HOST")
	var	DBPort, _  	= strconv.Atoi(os.Getenv("DB_PORT"))
	var	DBUser    	= os.Getenv("DB_USER")
	var	DBPassword 	= os.Getenv("DB_PASSWORD")
	var	DBName   	= os.Getenv("DB_NAME")

	// init database
	err = database.InitDB(DBHost, DBPort, DBUser, DBPassword, DBName)
	if err != nil {
		log.Fatal(err)
	}
	//testdata.CleanDB()
	database.SetupAllTables()
	//testdata.FillDB()
	defer database.CloseDB()

	// define routes
	r := mux.NewRouter()

	// serve frontend
	frontend := r.PathPrefix("/").Subrouter()
	frontend.Handle("/",  http.FileServer(getPWAFileSystem()))

	// serve api
	api := r.PathPrefix("/api/v1").Subrouter()
	charger := api.PathPrefix("/charger").Subrouter()
	charger.HandleFunc("/all", chargingStation.OverviewHandler).Methods(http.MethodGet)
	charger.HandleFunc("/{id}", chargingStation.DetailsHandler).Methods(http.MethodGet)
	charger.HandleFunc("/", chargingStation.CreateHandler).Methods(http.MethodPost)
	charger.HandleFunc("/", chargingStation.UpdateHandler).Methods(http.MethodPut)
	charger.HandleFunc("/{id}", chargingStation.DeleteHandler).Methods(http.MethodDelete)

	log.Fatalln(http.ListenAndServeTLS(":"+os.Getenv("SERVER_PORT"), os.Getenv("SERVER_CERT_FILE_PATH"), os.Getenv("SERVER_KEY_FILE_PATH"), r))
}

func getPWAFileSystem() http.FileSystem {
	fsys, err := fs.Sub(frontend.PWA, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}

func createDefaultDirs() {
	createDefaultConfigDirs()
	createCertsDir()
}

func createDefaultConfigDirs() {
	files, err := config.EnvFiles.ReadDir(".")
	if err != nil {
		log.Println(err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			fileContent, err := config.EnvFiles.ReadFile(file.Name())
			if err != nil {
				log.Println(err)
			}

			var fileName = strings.TrimSuffix(file.Name(), ".default")

			if strings.HasSuffix(fileName, ".env") {
				var filePath = filepath.Join(config.GetConfigFilePath(), fileName)
				if !helper.Exists(filePath) {
					helper.WriteFile(fileContent, filePath)
					fmt.Printf("Default file %s created\n", fileName)
				}
			}
		}
	}
}

func createCertsDir() {
	path := "certs"
	if !helper.Exists(path) {
		err := helper.CreateDir(path)
		if err != nil{
			log.Println(err)
		}
	}
}
