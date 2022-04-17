package currency

import (
	"PowerShare/database"
	"PowerShare/models"
	"log"
)

func CreateCurrency(currency models.Currency) {
	// WARNING: no checks used because this is only used internally and not exposed by handlers
	sqlInsertStatement := `INSERT INTO currencies (abbreviation, symbol) VALUES ($1, $2) RETURNING id`
	err := database.DB.QueryRow(sqlInsertStatement, currency.Abbreviation, currency.Symbol).Scan(&currency.ID)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
}
