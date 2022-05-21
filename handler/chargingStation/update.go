package chargingStation

import (
	"PowerShare/database"
	"PowerShare/helper/user"
	"PowerShare/models"
	"fmt"
)

func updateCharger(charger models.Charger, email string) (int64, error) {
	userId, err := user.GetId(email)
	if err != nil {
		return -1, fmt.Errorf("no user found. %v", err)
	}

	sqlStatement := `UPDATE chargers SET title=$2, position=POINT($3, $4), cost=$5, currencyId=(SELECT id FROM currencies WHERE abbreviation=$6), isOccupied=$7, description=$8, shelly_device_id=$9 WHERE id=$1 AND userid=$10 RETURNING id`

	var id int64
	err = database.DB.QueryRow(sqlStatement, charger.ID, charger.Title, charger.Position.Lat, charger.Position.Lng, charger.Cost.Amount, charger.Cost.Currency.Abbreviation, charger.IsOccupied, charger.Description, charger.TechnicalData.ShellyDeviceId, userId).Scan(&id)

	return id, err
}
