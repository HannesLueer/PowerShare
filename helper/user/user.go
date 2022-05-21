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
