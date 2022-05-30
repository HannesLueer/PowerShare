package user

import (
	"PowerShare/database"
	"fmt"
	"log"
)

func GetId(email string) (id int64, err error) {
	sqlSelectStatement := `SELECT id FROM users WHERE email=$1`
	err = database.DB.QueryRow(sqlSelectStatement, email).Scan(&id)
	if err != nil {
		log.Println(err)
		return -1, fmt.Errorf("internal error")
	}
	return id, nil
}

func GetPayoutEmail(id int64) (email string, err error) {
	sqlSelectStatement := `SELECT paypal_email FROM users WHERE id=$1`
	err = database.DB.QueryRow(sqlSelectStatement, id).Scan(&email)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("internal error")
	}
	return email, nil
}
