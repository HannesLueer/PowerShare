package chargingStation

import (
	"PowerShare/database"
	"PowerShare/models"
)

func updateCharger(charger models.Charger) (int64, error){
	sqlStatement := `UPDATE chargers SET title=$2, position=POINT($3, $4), cost=$5, isOccupied=$6 WHERE id=$1 RETURNING id`

	var id int64
	err := database.DB.QueryRow(sqlStatement, charger.ID, charger.Title, charger.Position.Lat, charger.Position.Lng, charger.Cost, charger.IsOccupied).Scan(&id)

	return id, err
}
