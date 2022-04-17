package main

import (
	"PowerShare/config"
	"PowerShare/database"
	"PowerShare/frontend"
	"PowerShare/handler/chargingStation"
	"PowerShare/handler/currency"
	"PowerShare/handler/user"
	"PowerShare/helper"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// define mime types
	err := mime.AddExtensionType(".js", "text/javascript")
	if err != nil {
		log.Println(err)
	}
	err = mime.AddExtensionType(".css", "text/css")
	if err != nil {
		log.Println(err)
	}

	// create default directories and files
	createDefaultDirs()

	// load .env files
	err = godotenv.Load(
		filepath.Join(config.GetConfigFilePath(), "db.env"),
		filepath.Join(config.GetConfigFilePath(), "server.env"),
	)
	if err != nil {
		log.Fatal(err)
	}

	var DBHost = os.Getenv("DB_HOST")
	var DBPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	var DBUser = os.Getenv("DB_USER")
	var DBPassword = os.Getenv("DB_PASSWORD")
	var DBName = os.Getenv("DB_NAME")

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

	// serve api
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	chargerRouter := apiRouter.PathPrefix("/charger").Subrouter()
	chargerRouter.HandleFunc("/all", chargingStation.OverviewHandler).Methods(http.MethodGet)
	chargerRouter.HandleFunc("/my", user.IsAuthorized(chargingStation.OverviewOwnHandler)).Methods(http.MethodGet)
	chargerRouter.HandleFunc("/{id}", chargingStation.DetailsHandler).Methods(http.MethodGet)
	chargerRouter.HandleFunc("/", user.IsAuthorized(chargingStation.CreateHandler)).Methods(http.MethodPost)
	chargerRouter.HandleFunc("/", user.IsAuthorized(chargingStation.UpdateHandler)).Methods(http.MethodPut)
	chargerRouter.HandleFunc("/{id}", user.IsAuthorized(chargingStation.DeleteHandler)).Methods(http.MethodDelete)
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/signup", user.SignUpHandler).Methods(http.MethodPost)
	userRouter.HandleFunc("/signin", user.SignInHandler).Methods(http.MethodPost)
	userRouter.HandleFunc("/", user.IsAuthorized(user.UpdateHandler)).Methods(http.MethodPut)
	userRouter.HandleFunc("/", user.IsAuthorized(user.DeleteHandler)).Methods(http.MethodDelete)
	userRouter.HandleFunc("/", user.IsAuthorized(user.GetHandler)).Methods(http.MethodGet)
	currencyRouter := apiRouter.PathPrefix("/currency").Subrouter()
	currencyRouter.HandleFunc("/all", currency.OverviewHandler).Methods(http.MethodGet)

	// serve frontend
	frontendRouter := r.PathPrefix("/").Subrouter()
	frontendRouter.PathPrefix("/").HandlerFunc(frontend.SpaHandler)

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*://localhost:3000"},
		AllowCredentials: true,
		Debug:            false,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPost},
	})
	corsMux := c.Handler(r)

	log.Fatalln(http.ListenAndServeTLS(":"+os.Getenv("SERVER_PORT"), os.Getenv("SERVER_CERT_FILE_PATH"), os.Getenv("SERVER_KEY_FILE_PATH"), corsMux))
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
					err := helper.WriteFile(fileContent, filePath)
					if err != nil {
						log.Println(err)
					} else {
						log.Printf("Default file %s created", fileName)
					}
				}
			}
		}
	}
}

func createCertsDir() {
	path := "certs"
	if !helper.Exists(path) {
		err := helper.CreateDir(path)
		if err != nil {
			log.Println(err)
		}
	}
}
