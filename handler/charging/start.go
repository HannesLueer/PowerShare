package charging

import (
	"PowerShare/database"
	"PowerShare/helper/charging"
	"PowerShare/helper/gocardless"
	"PowerShare/helper/jwt"
	"PowerShare/helper/shelly"
	"PowerShare/helper/smartme"
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

	// get user email
	tokenStr, errCode, err := jwt.GetToken(r)
	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}
	email, err := jwt.GetEmailFromToken(tokenStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to read token", http.StatusBadRequest)
		return
	}

	// start charging
	httpErrorCode, err := startCharging(chargerId, email)
	if err != nil {
		http.Error(w, err.Error(), httpErrorCode)
		return
	}

	return
}

func startCharging(chargerID int64, userEmail string) (httpErrorCode int, error error) {
	// check mandate of user
	_, err := gocardless.GetMandateIdFromEmail(userEmail)
	if err != nil {
		return http.StatusForbidden, fmt.Errorf("no mandate was found")
	}

	// check if charging station is available
	isChargerAvailable, err := charging.IsChargerAvailable(chargerID)
	if err != nil || !isChargerAvailable {
		return http.StatusBadRequest, fmt.Errorf("charger is not available")
	}

	// read electricity meter
	counterReadingKWH, err := smartme.ReadCounter(userEmail, chargerID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("an error occurred during communication with the meter")
	}

	// turn power on
	statusCode, err := charging.SwitchPower(chargerID, shelly.On)
	if err != nil {
		return statusCode, err
	}

	// write loading process in db
	writeChargingProcessDB(userEmail, chargerID, counterReadingKWH)

	return http.StatusOK, nil
}

func writeChargingProcessDB(userEmail string, chargerID int64, meterStartCount float64) (httpErrorCode int, error error) {
	// add charging process
	sqlStatement := `INSERT INTO charging_processes (user_id, charger_id, meter_start_count) VALUES (((SELECT id FROM users WHERE email=$1)), $2, $3) RETURNING id`
	var id int64
	err := database.DB.QueryRow(sqlStatement, userEmail, chargerID, meterStartCount).Scan(&id)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}

	// set charger occupied
	err = charging.UpdateChargerAvailability(chargerID, false)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}

	return http.StatusOK, nil
}
