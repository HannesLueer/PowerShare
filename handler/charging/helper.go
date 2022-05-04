package charging

import (
	"PowerShare/database"
)

func isChargerAvailable(chargerID int64) (isAvailable bool, err error) {
	sqlStatement := `SELECT isoccupied FROM chargers WHERE id=$1`
	err = database.DB.QueryRow(sqlStatement, chargerID).Scan(&isAvailable)
	return !isAvailable, err
}
