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
	sqlStatement := `CREATE TABLE chargers(id serial, title text, position point, cost numeric, currencyId integer, isOccupied boolean, description text, shelly_device_id integer, userId integer, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func createChargingProcessesTable() {
	sqlStatement := `CREATE TABLE charging_processes(id serial, userId integer, chargerId integer, paypal_order_id text, amount numeric, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func createShellyDevicesTable() {
	sqlStatement := `CREATE TABLE shelly_devices(device_id integer, user_id integer, device_type text, device_code text, host text, PRIMARY KEY( device_id ))`
	DB.Exec(sqlStatement)
}

func SetupAllTables() {
	createCurrenciesTable()
	createUsersTable()
	createChargersTable()
	createChargingProcessesTable()
	createShellyDevicesTable()

	log.Println("DB tables created")
}
