package chargingStation

import (
	"PowerShare/database"
	"PowerShare/models"
	"regexp"
	"strconv"
)

func getAllCharger() ([]models.Charger, error) {
	rows, err := database.DB.Query("SELECT id, title, position, cost, isoccupied FROM chargers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chargers []models.Charger

	for rows.Next() {
		var charger models.Charger
		var coordinateStr string

		err := rows.Scan(&charger.ID, &charger.Title, &coordinateStr, &charger.Cost, &charger.IsOccupied)
		if err != nil {
			return nil, err
		}

		charger.Position, err = getCoordinateFromString(coordinateStr)
		if err != nil {
			return nil, err
		}

		chargers = append(chargers, charger)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return chargers, nil
}

func getCharger(id int64) (models.Charger, error) {
	sqlStatement := `SELECT id, title, position, cost, isoccupied FROM chargers WHERE id=$1`

	var charger models.Charger
	var coordinateStr string
	err := database.DB.QueryRow(sqlStatement, id).Scan(&charger.ID, &charger.Title, &coordinateStr, &charger.Cost, &charger.IsOccupied)
	if err != nil {
		return charger, err
	}

	charger.Position, err = getCoordinateFromString(coordinateStr)
	if err != nil {
		return charger, err
	}

	return charger, err
}

func getChargersOfUser(email string) {

}

func getCoordinateFromString(coordinate string) (models.Coordinate, error) {
	t := regexp.MustCompile(`\(|\)|,| `)
	array := t.Split(coordinate, -1)
	lat, err := strconv.ParseFloat(array[2], 64)
	lng, err := strconv.ParseFloat(array[1], 64)
	return models.Coordinate{Lat: lat, Lng: lng}, err
}
