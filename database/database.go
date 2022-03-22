package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL
)

var DB *sql.DB

// InitDB sets up the connection pool global variable.
func InitDB(DBHost string, DBPort int, DBUser string, DBPassword string, DBName string) error {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	return DB.Ping()
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
