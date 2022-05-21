package charging

import (
	"PowerShare/database"
	"PowerShare/helper/charger"
	"PowerShare/helper/shelly"
	"PowerShare/models"
)

func IsChargerAvailable(chargerID int64) (isAvailable bool, err error) {
	sqlStatement := `SELECT isoccupied FROM chargers WHERE id=$1`
	err = database.DB.QueryRow(sqlStatement, chargerID).Scan(&isAvailable)
	return !isAvailable, err
}

func UpdateChargerAvailability(chargerID int64, isAvailable bool) (err error) {
	sqlStatement := `UPDATE chargers SET isoccupied=$2 WHERE id=$1 RETURNING id`
	err = database.DB.QueryRow(sqlStatement, chargerID, isAvailable).Scan(&chargerID)
	return err
}

func GetElectricityAmount(chargerID int64) (amountKWH float64, err error) {
	//TODO return the amount of electricity measured by the electric meter of the charger

	return 12, nil
}

func SwitchPower(chargerID int64, mode shelly.Mode) (httpStatusCode int, err error) {
	shellyDeviceId := charger.GetShellyDeviceID(chargerID)
	return shelly.TurnPower(shellyDeviceId, mode)
}

func GetCostPerKWH(chargerID int64) (cost models.Cost, err error) {
	sqlStatement := `SELECT chargers.cost, currencies.abbreviation, currencies.symbol FROM chargers INNER JOIN currencies ON chargers.currencyId=currencies.id WHERE chargers.id=$1`
	err = database.DB.QueryRow(sqlStatement, chargerID).Scan(&cost.Amount, cost.Currency.Abbreviation, cost.Currency.Symbol)
	return cost, err
}
