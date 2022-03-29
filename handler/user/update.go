package user

import (
	"PowerShare/database"
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header["Token"] == nil {
		http.Error(w, "no token found", http.StatusBadRequest)
		return
	}

	var tokenStr = r.Header["Token"][0]
	var oldEmail, err = getEmailFromToken(tokenStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to read token", http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "error in reading body", http.StatusBadRequest)
		return
	}

	_, errCode, err := update(oldEmail, user)
	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}

	// sign in
	token, errCode, err := signIn(models.Authentication{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}

	json.NewEncoder(w).Encode(token)
	return
}

func update(oldEmail string, user models.User) (id int64, errCode int, error error) {
	//checks if email is already in use
	var dbUser models.User
	sqlSelectStatement := `SELECT * FROM users WHERE email=$1`
	err := database.DB.QueryRow(sqlSelectStatement, user.Email).Scan(&dbUser.ID, &dbUser.Name, &dbUser.Email, &dbUser.Password, &dbUser.Role)
	if err != nil {
		log.Println(err)
	}
	if dbUser.Email != "" {
		return -1, http.StatusBadRequest, fmt.Errorf("email already in use")
	}

	// hash password
	user.Password, error = generateHashPassword(user.Password)
	if error != nil {
		log.Println("error in password hash")
		return -1, http.StatusInternalServerError, fmt.Errorf("internal error")
	}

	// update user in database
	sqlUpdateStatement := `UPDATE users SET name=$2, email=$3, password=$4 WHERE email=$1 RETURNING id`
	error = database.DB.QueryRow(sqlUpdateStatement, oldEmail, user.Name, user.Email, user.Password).Scan(&id)
	if error != nil {
		log.Println(error)
		return id, http.StatusInternalServerError, fmt.Errorf("internal error")
	}
	return id, http.StatusOK, nil
}
