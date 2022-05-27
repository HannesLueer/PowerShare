package gocardless

import (
	"PowerShare/helper/gocardless"
	"PowerShare/helper/jwt"
	"log"
	"net/http"
)

func NewMandateHandler(w http.ResponseWriter, r *http.Request) {
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

	url, err := gocardless.NewMandate(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(url))

	return
}
