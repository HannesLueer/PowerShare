package chargingStation

import (
	"PowerShare/database"
	"PowerShare/handler/user"
	"PowerShare/models"
	"fmt"
	"regexp"
	"strconv"
)

func getAllCharger() ([]models.Charger, error) {
	rows, err := database.DB.Query("SELECT chargers.id, chargers.title, chargers.position, chargers.cost, currencies.abbreviation, currencies.symbol, chargers.isoccupied, chargers.description FROM chargers INNER JOIN currencies ON chargers.currencyId=currencies.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chargers []models.Charger

	for rows.Next() {
		var charger models.Charger
		var coordinateStr string

		err := rows.Scan(&charger.ID, &charger.Title, &coordinateStr, &charger.Cost.Amount, &charger.Cost.Currency.Abbreviation, &charger.Cost.Currency.Symbol, &charger.IsOccupied, &charger.Description)
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
	sqlStatement := `SELECT chargers.id, chargers.title, chargers.position, chargers.cost, currencies.abbreviation, currencies.symbol, chargers.isoccupied, chargers.description FROM chargers INNER JOIN currencies ON chargers.currencyId=currencies.id WHERE chargers.id=$1`

	var charger models.Charger
	var coordinateStr string
	err := database.DB.QueryRow(sqlStatement, id).Scan(&charger.ID, &charger.Title, &coordinateStr, &charger.Cost.Amount, &charger.Cost.Currency.Abbreviation, &charger.Cost.Currency.Symbol, &charger.IsOccupied, &charger.Description)
	if err != nil {
		return charger, err
	}

	charger.Position, err = getCoordinateFromString(coordinateStr)
	if err != nil {
		return charger, err
	}

	return charger, err
}

func getChargersOfUser(email string) ([]models.Charger, error) {
	userId, err := user.GetId(email)
	if err != nil {
		return nil, fmt.Errorf("no user found. %v", err)
	}

	rows, err := database.DB.Query("SELECT chargers.id, chargers.title, chargers.position, chargers.cost, currencies.abbreviation, currencies.symbol, chargers.isoccupied, chargers.description FROM chargers INNER JOIN currencies ON chargers.currencyId=currencies.id WHERE chargers.userid=$1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chargers []models.Charger

	for rows.Next() {
		var charger models.Charger
		var coordinateStr string

		err := rows.Scan(&charger.ID, &charger.Title, &coordinateStr, &charger.Cost.Amount, &charger.Cost.Currency.Abbreviation, &charger.Cost.Currency.Symbol, &charger.IsOccupied, &charger.Description)
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

func getCoordinateFromString(coordinate string) (models.Coordinate, error) {
	t := regexp.MustCompile(`\(|\)|,| `)
	array := t.Split(coordinate, -1)
	lat, err := strconv.ParseFloat(array[1], 64)
	lng, err := strconv.ParseFloat(array[2], 64)
	return models.Coordinate{Lat: lat, Lng: lng}, err
}
