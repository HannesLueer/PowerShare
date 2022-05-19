package charger

import (
	"PowerShare/database"
	"log"
)

func GetUserId(chargerID int64) (userID int64) {
	sqlStatement := `SELECT chargers.userid FROM chargers WHERE chargers.id=$1`
	err := database.DB.QueryRow(sqlStatement, chargerID).Scan(&userID)
	if err != nil {
		userID = -1
		log.Printf("Unable to execute the query. %v", err)
	}
	return userID
}
