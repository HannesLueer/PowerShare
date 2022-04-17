package currency

import (
	"PowerShare/database"
	"PowerShare/models"
)

func getCurrencies() ([]models.Currency, error) {
	rows, err := database.DB.Query("SELECT id, abbreviation, symbol FROM currencies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []models.Currency

	for rows.Next() {
		var currency models.Currency

		err := rows.Scan(&currency.ID, &currency.Abbreviation, &currency.Symbol)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return currencies, nil
}
