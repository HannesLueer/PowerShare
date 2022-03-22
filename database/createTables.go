package database

import "fmt"

func createChargersTable(){
	sqlStatement := `CREATE TABLE chargers(id serial, title text, position point, cost numeric, isOccupied boolean, PRIMARY KEY( id ))`
	DB.Exec(sqlStatement)
}

func SetupAllTables(){
	createChargersTable()

	fmt.Println("DB tables created")
}
