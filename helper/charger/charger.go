package charger

import (
	"PowerShare/database"
	"PowerShare/models"
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

func GetShellyDeviceID(chargerID int64) (shellyDeviceId int64) {
	sqlStatement := `SELECT chargers.shelly_device_id FROM chargers WHERE chargers.id=$1`
	err := database.DB.QueryRow(sqlStatement, chargerID).Scan(&shellyDeviceId)
	if err != nil {
		shellyDeviceId = -1
		log.Printf("Unable to execute the query. %v", err)
	}
	return shellyDeviceId
}

func GetSmartmeSerialNumber(chargerID int64) (smartmeSerialNumber string) {
	sqlStatement := `SELECT smartme_serial_number FROM chargers WHERE chargers.id=$1`
	err := database.DB.QueryRow(sqlStatement, chargerID).Scan(&smartmeSerialNumber)
	if err != nil {
		smartmeSerialNumber = ""
		log.Printf("Unable to execute the query. %v", err)
	}
	return smartmeSerialNumber
}

func IsChargerValid(charger models.Charger) bool {
	return charger.Title != "" &&
		isLngValid(charger.Position.Lng) &&
		isLatValid(charger.Position.Lat) &&
		charger.TechnicalData.SmartmeSerialNumber != "" &&
		charger.TechnicalData.ShellyDeviceId >= 0 &&
		charger.Cost.Amount > 0 &&
		charger.Cost.Currency.Abbreviation != ""
}

func isLatValid(lat float64) bool {
	return lat >= -90 && lat <= 90
}

func isLngValid(lng float64) bool {
	return lng >= -180 && lng <= 180
}
