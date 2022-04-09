package user

import (
	"PowerShare/database"
	"fmt"
	"log"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header["Token"] == nil {
		http.Error(w, "No token found", http.StatusBadRequest)
		return
	}

	var tokenStr = r.Header["Token"][0]
	var email, err = getEmailFromToken(tokenStr)

	if err != nil {
		log.Println(err)
		http.Error(w, "unable to read token", http.StatusBadRequest)
		return
	}

	errCode, err := deleteUser(email)
	if err != nil {
		http.Error(w, err.Error(), errCode)
	}
	return
}

func deleteUser(email string) (errCode int, error error) {
	sqlStatement := `DELETE FROM users WHERE email = $1`

	_, err := database.DB.Exec(sqlStatement, email)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}

	return http.StatusOK, nil
}
