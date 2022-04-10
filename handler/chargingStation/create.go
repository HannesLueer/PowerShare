package chargingStation

import (
	"PowerShare/database"
	"PowerShare/handler/user"
	"PowerShare/models"
	"log"
)

func CreateCharger(charger models.Charger, userEmail string) int64 {
	userId, err := user.GetId(userEmail)
	if err != nil {
		log.Printf("No user found. %v", err)
		return -1
	}

	sqlStatement := `INSERT INTO chargers (title, position, cost, isOccupied, userId) VALUES ($1, POINT($2, $3), $4, $5, $6) RETURNING id`

	var id int64
	err = database.DB.QueryRow(sqlStatement, charger.Title, charger.Position.Lat, charger.Position.Lng, charger.Cost, charger.IsOccupied, userId).Scan(&id)

	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return -1
	}

	return id
}
