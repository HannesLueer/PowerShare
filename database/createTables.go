package database

import (
	"log"
)

func createCurrenciesTable() {
	sqlStatement := `CREATE TABLE currencies(id serial, abbreviation text, symbol text, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func createUsersTable() {
	sqlStatement := `CREATE TABLE users (id serial, name text, email text, password text, role int, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func createChargersTable() {
	sqlStatement := `CREATE TABLE chargers(id serial, title text, position point, cost numeric, currencyId integer, isOccupied boolean, userId integer, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func SetupAllTables() {
	createCurrenciesTable()
	createUsersTable()
	createChargersTable()

	log.Println("DB tables created")
}
