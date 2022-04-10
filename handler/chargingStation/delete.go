package chargingStation

import (
	"PowerShare/database"
	"PowerShare/handler/user"
	"fmt"
	"log"
)

func deleteCharger(id int64, email string) error {
	userId, err := user.GetId(email)
	if err != nil {
		return fmt.Errorf("no user found. %v", err)
	}

	sqlStatement := `DELETE FROM chargers WHERE id = $1 AND userid=$2`

	_, err = database.DB.Exec(sqlStatement, id, userId)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
	}

	return err
}
