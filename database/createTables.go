package database

import (
	"log"
)

func createUsersTable() {
	sqlStatement := `CREATE TABLE users (id serial, name text, email text, password text, role int, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func createChargersTable() {
	sqlStatement := `CREATE TABLE chargers(id serial, title text, position point, cost numeric, isOccupied boolean, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func SetupAllTables() {
	createUsersTable()
	createChargersTable()

	log.Println("DB tables created")
}
