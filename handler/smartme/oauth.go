package smartme

import (
	"PowerShare/helper/jwt"
	"PowerShare/helper/smartme"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func AuthorizationCodeHandler(w http.ResponseWriter, r *http.Request) {
	// get email
	tokenStr, errCode, err := jwt.GetToken(r)
	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}
	userEmail, err := jwt.GetEmailFromToken(tokenStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to read token", http.StatusBadRequest)
		return
	}

	// read code from url
	vars := mux.Vars(r)
	authCode := vars["code"]

	// get access token and save it
	err = smartme.RequestAndSaveFirstAccessToken(userEmail, authCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}
