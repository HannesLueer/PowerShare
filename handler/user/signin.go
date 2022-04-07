package user

import (
	"PowerShare/database"
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var authDetails models.Authentication
	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		json.NewEncoder(w).Encode("Error in reading body")
		return
	}

	token, errCode, err := signIn(authDetails)

	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}
	json.NewEncoder(w).Encode(token)
	return
}

func signIn(authDetails models.Authentication) (token models.Token, httpErrorCode int, error error) {
	var authUser models.User
	var t models.Token
	sqlSelectStatement := `SELECT * FROM users WHERE email=$1`
	err := database.DB.QueryRow(sqlSelectStatement, authDetails.Email).Scan(&authUser.ID, &authUser.Name, &authUser.Email, &authUser.Password, &authUser.Role)
	if err != nil {
		log.Println(err)
	}
	if authUser.Email == "" {
		return t, http.StatusBadRequest, fmt.Errorf("username or password is incorrect")
	}

	check := checkPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		return t, http.StatusBadRequest, fmt.Errorf("username or password is incorrect")
	}

	validToken, err := generateJWT(authUser.Email, authUser.Role)
	if err != nil {
		return t, http.StatusInternalServerError, fmt.Errorf("failed to generate token")
	}

	t.Email = authUser.Email
	t.Role = authUser.Role
	t.TokenString = validToken
	return t, http.StatusOK, nil
}