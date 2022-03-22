package chargingStation

import (
	"PowerShare/database"
	"PowerShare/models"
	"log"
)

func CreateCharger(charger models.Charger) int64{
	sqlStatement := `INSERT INTO chargers (title, position, cost, isOccupied) VALUES ($1, POINT($2, $3), $4, $5) RETURNING id`

	var id int64
	err := database.DB.QueryRow(sqlStatement, charger.Title, charger.Position.Lat, charger.Position.Lng, charger.Cost, charger.IsOccupied).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return -1
	}

	return id
}
