package chargingStation

import (
	"PowerShare/database"
	"PowerShare/handler/user"
	"PowerShare/models"
	"fmt"
)

func updateCharger(charger models.Charger, email string) (int64, error) {
	userId, err := user.GetId(email)
	if err != nil {
		return -1, fmt.Errorf("no user found. %v", err)
	}

	sqlStatement := `UPDATE chargers SET title=$2, position=POINT($3, $4), cost=$5, isOccupied=$6 WHERE id=$1 AND userid=$7 RETURNING id`

	var id int64
	err = database.DB.QueryRow(sqlStatement, charger.ID, charger.Title, charger.Position.Lat, charger.Position.Lng, charger.Cost, charger.IsOccupied, userId).Scan(&id)

	return id, err
}
