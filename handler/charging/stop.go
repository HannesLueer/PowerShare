package charging

import (
	"PowerShare/database"
	"PowerShare/handler/charging/helper"
	"PowerShare/handler/user"
	"PowerShare/models"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func StopHandler(w http.ResponseWriter, r *http.Request) {
	// read chargerId from url
	vars := mux.Vars(r)
	chargerIdStr := vars["chargerId"]
	chargerId, err := strconv.ParseInt(chargerIdStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	// stop charging
	httpErrorCode, err := stopCharging(chargerId, email)
	if err != nil {
		http.Error(w, err.Error(), httpErrorCode)
		return
	}
	return
}

func stopCharging(chargerID int64, userEmail string) (httpErrorCode int, error error) {
	// turn power off
	err := helper.SwitchPower(chargerID, false)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// read electric meter
	chargedEnergyKWH, err := helper.GetElectricityAmount(chargerID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// calculate cost
	cost, err := getCost(chargerID, chargedEnergyKWH)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// update payment
	paypalOrderID, err := getPaypalOrderID(chargerID, userEmail)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = helper.UpdateOrderPaypal(paypalOrderID, cost)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = helper.CaptureFundsPaypal(paypalOrderID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// update database
	httpErrorCode, err = writeAmountDB(userEmail, chargerID, chargedEnergyKWH)
	if err != nil {
		return httpErrorCode, err
	}

	httpErrorCode, err = setChargerAvailable(chargerID)
	if err != nil {
		return httpErrorCode, err
	}

	return http.StatusOK, nil
}

func writeAmountDB(userEmail string, chargerID int64, amount float64) (httpErrorCode int, error error) {
	sqlStatement := `UPDATE charging_processes SET amount=$3 WHERE (chargerid=$1 AND userid=(SELECT id FROM users WHERE email=$2) AND amount IS NULL)`
	var id int64
	err := database.DB.QueryRow(sqlStatement, chargerID, userEmail, amount).Scan(&id)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}
	return http.StatusOK, nil
}

func setChargerAvailable(chargerID int64) (httpErrorCode int, error error) {
	err := helper.UpdateChargerAvailability(chargerID, true)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}
	return http.StatusOK, nil
}

func getPaypalOrderID(chargerID int64, userEmail string) (paypalOrderId string, err error) {
	sqlStatement := `SELECT paypal_order_id FROM charging_processes WHERE (chargerid=$1 AND userid=(SELECT id FROM users WHERE email=$2) AND amount IS NULL)`
	err = database.DB.QueryRow(sqlStatement, chargerID, userEmail).Scan(&paypalOrderId)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return "", fmt.Errorf("internal error")
	}
	return paypalOrderId, nil
}

func getCost(chargerID int64, chargedEnergyKWH float64) (cost models.Cost, err error) {
	costPerKWH, err := helper.GetCostPerKWH(chargerID)
	if err != nil {
		return models.Cost{}, fmt.Errorf("internal error")
	}

	return models.Cost{
		Amount:   float32(float64(costPerKWH.Amount) * chargedEnergyKWH),
		Currency: costPerKWH.Currency,
	}, nil
}
