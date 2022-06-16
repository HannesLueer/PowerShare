package user

import (
	"PowerShare/database"
	"PowerShare/helper/jwt"
	userHelper "PowerShare/helper/user"
	"PowerShare/models"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "error in reading body", http.StatusBadRequest)
		return
	}

	if !userHelper.IsUserValid(user) {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	_, errCode, err := SignUp(user)
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

func SignUp(user models.User) (id int64, httpErrorCode int, error error) {
	//checks if email is already register or not
	var dbUser models.User
	sqlSelectStatement := `SELECT * FROM users WHERE email=$1`
	err := database.DB.QueryRow(sqlSelectStatement, user.Email).Scan(&dbUser.ID, &dbUser.Name, &dbUser.Email, &dbUser.PasswordHash, &dbUser.Role)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
	}
	if dbUser.Email != "" {
		return -1, http.StatusBadRequest, fmt.Errorf("email already in use")
	}

	user.PasswordHash, err = jwt.GenerateHashPassword(user.Password)
	if err != nil {
		log.Printf("error in password hash: %v", err)
		return -1, http.StatusInternalServerError, fmt.Errorf("internal error")
	}

	//insert user details in database
	sqlInsertStatement := `INSERT INTO users (name, email, paypal_email, password, role) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = database.DB.QueryRow(sqlInsertStatement, user.Name, user.Email, user.PaypalEmail, user.PasswordHash, user.Role).Scan(&id)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return -1, http.StatusInternalServerError, fmt.Errorf("internal error")
	}
	return id, http.StatusOK, nil
}
