package chargingStation

import (
	"PowerShare/database"
	"PowerShare/helper/user"
	"PowerShare/models"
	"log"
)

func CreateCharger(charger models.Charger, userEmail string) int64 {
	userId, err := user.GetId(userEmail)
	if err != nil {
		log.Printf("No user found. %v", err)
		return -1
	}

	sqlStatement := `INSERT INTO chargers (title, position, cost, currencyId, isOccupied, description, shelly_device_id, smartme_serial_number, userId) VALUES ($1, POINT($2, $3), $4, (SELECT id FROM currencies WHERE abbreviation=$5), $6, $7, $8, $9, $10) RETURNING id`

	var id int64
	err = database.DB.QueryRow(sqlStatement, charger.Title, charger.Position.Lat, charger.Position.Lng, charger.Cost.Amount, charger.Cost.Currency.Abbreviation, charger.IsOccupied, charger.Description, charger.TechnicalData.ShellyDeviceId, charger.TechnicalData.SmartmeSerialNumber, userId).Scan(&id)

	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return -1
	}

	return id
}
