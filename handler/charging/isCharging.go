package charging

import (
	"PowerShare/database"
	"PowerShare/helper/jwt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func IsUserChargingHandler(w http.ResponseWriter, r *http.Request) {
	// read chargerId values from url
	vars := mux.Vars(r)
	chargerIdStr := vars["chargerId"]
	chargerId, err := strconv.ParseInt(chargerIdStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get user email
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

	isUserChargingAtCharger(userEmail, chargerId)

	w.Write([]byte("false"))
}

func isUserChargingAtCharger(userEmail string, chargerId int64) bool {
	sqlStatement := `SELECT id FROM charging_processes WHERE (charger_id=$1 AND user_id=(SELECT id FROM users WHERE email=$2) AND amount IS NULL)`
	var id int64 = -1
	database.DB.QueryRow(sqlStatement, chargerId, userEmail).Scan(&id)

	return !(id == -1)
}
