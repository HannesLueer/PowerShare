package user

import (
	"PowerShare/database"
	"PowerShare/helper/jwt"
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header["Token"] == nil {
		http.Error(w, "no token found", http.StatusBadRequest)
		return
	}

	var tokenStr = r.Header["Token"][0]
	var email, err = jwt.GetEmailFromToken(tokenStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to read token", http.StatusBadRequest)
		return
	}

	user, errCode, err := get(email)
	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}

	json.NewEncoder(w).Encode(user)
	return
}

func get(email string) (user models.User, errCode int, error error) {
	sqlSelectStatement := `SELECT name, email, paypal_email, role FROM users WHERE email=$1`
	err := database.DB.QueryRow(sqlSelectStatement, email).Scan(&user.Name, &user.Email, &user.PaypalEmail, &user.Role)
	if err != nil {
		log.Println(err)
		return models.User{}, http.StatusInternalServerError, fmt.Errorf("internal error")
	}

	return user, http.StatusOK, nil
}
