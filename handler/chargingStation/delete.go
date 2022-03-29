package chargingStation

import (
	"PowerShare/database"
	"log"
)

func deleteCharger(id int64) error {
	sqlStatement := `DELETE FROM chargers WHERE id = $1`

	_, err := database.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
	}

	return err
}
