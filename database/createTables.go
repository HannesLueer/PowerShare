package database

import (
	"log"
)

func createCurrenciesTable() {
	sqlStatement := `CREATE TABLE currencies(id serial, abbreviation text, symbol text, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func createUsersTable() {
	sqlStatement := `CREATE TABLE users (id serial, name text, email text, paypal_email text, password text, role int, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func createChargersTable() {
	sqlStatement := `CREATE TABLE chargers(id serial, title text, position point, cost numeric, currencyId integer, isOccupied boolean, description text, shelly_device_id integer, smartme_serial_number text, userId integer, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func createChargingProcessesTable() {
	sqlStatement := `CREATE TABLE charging_processes(id serial, user_id integer, charger_id integer, payment_id text, payout_id text, meter_start_count numeric, amount numeric, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func createShellyDevicesTable() {
	sqlStatement := `CREATE TABLE shelly_devices(device_id integer, user_id integer, device_type text, device_code text, host text, PRIMARY KEY( device_id ))`
	DB.Exec(sqlStatement)
}

func createSmartmeTokensTable() {
	sqlStatement := `CREATE TABLE smartme_tokens(user_id integer, access_token text, token_type text, refresh_token text, PRIMARY KEY( user_id ))`
	DB.Exec(sqlStatement)
}

func SetupAllTables() {
	createCurrenciesTable()
	createUsersTable()
	createChargersTable()
	createChargingProcessesTable()
	createShellyDevicesTable()
	createSmartmeTokensTable()

	log.Println("DB tables created")
}
