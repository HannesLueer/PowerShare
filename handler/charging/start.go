package charging

import (
	"PowerShare/database"
	"PowerShare/handler/user"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func StartHandler(w http.ResponseWriter, r *http.Request) {
	// read values from url
	vars := mux.Vars(r)
	chargerIdStr := vars["chargerId"]
	chargerId, err := strconv.ParseInt(chargerIdStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paypalOrderID := vars["paypalOrderID"]

	// get user email
	tokenStr, errCode, err := user.GetToken(r)
	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}
	email, err := user.GetEmailFromToken(tokenStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to read token", http.StatusBadRequest)
		return
	}

	// start charging
	httpErrorCode, err := startCharging(chargerId, email, paypalOrderID)
	if err != nil {
		http.Error(w, err.Error(), httpErrorCode)
		return
	}
	return
}

func startCharging(chargerID int64, userEmail string, paypalOrderID string) (httpErrorCode int, error error) {
	// check if charging station is available
	isChargerAvailable, err := isChargerAvailable(chargerID)
	if err != nil || !isChargerAvailable {
		return http.StatusBadRequest, fmt.Errorf("charger is not available")
	}

	// TODO activate / read electricity meter

	// TODO turn on power

	// TODO change isAvailable to false

	// write loading process in db
	writeChargingProcessDB(userEmail, chargerID, paypalOrderID)

	return http.StatusOK, nil
}

func writeChargingProcessDB(userEmail string, chargerID int64, paypalOrderID string) (httpErrorCode int, error error) {
	sqlStatement := `INSERT INTO charging_processes (userid, chargerid, paypal_order_id) VALUES (((SELECT id FROM users WHERE email=$1)), $2, $3) RETURNING id`
	var id int64
	err := database.DB.QueryRow(sqlStatement, userEmail, chargerID, paypalOrderID).Scan(&id)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}
	return http.StatusOK, nil
}
