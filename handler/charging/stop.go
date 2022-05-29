package charging

import (
	"PowerShare/database"
	"PowerShare/helper/charger"
	"PowerShare/helper/charging"
	"PowerShare/helper/gocardless"
	"PowerShare/helper/jwt"
	"PowerShare/helper/paypal"
	"PowerShare/helper/shelly"
	"PowerShare/helper/user"
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
	statusCode, err := charging.SwitchPower(chargerID, shelly.Off)
	if err != nil {
		return statusCode, err
	}

	// read electric meter
	chargedEnergyKWH, err := charging.GetElectricityAmount(userEmail, chargerID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// calculate cost
	cost, err := getCost(chargerID, chargedEnergyKWH)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// create payment
	mandateId, err := gocardless.GetMandateIdFromEmail(userEmail)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	paymentId, err := gocardless.CreatingPaymentV2(cost, mandateId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// create payout
	ownerEmail, err := user.GetEmail(charger.GetUserId(chargerID))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	payoutResp, err := paypal.CreatePayoutPaypal(ownerEmail, getPayoutAmount(cost))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// update database
	httpErrorCode, err = writeAmountDB(userEmail, chargerID, chargedEnergyKWH, paymentId, payoutResp.BatchHeader.PayoutBatchId)
	if err != nil {
		return httpErrorCode, err
	}

	httpErrorCode, err = setChargerAvailable(chargerID)
	if err != nil {
		return httpErrorCode, err
	}

	return http.StatusOK, nil
}

func writeAmountDB(userEmail string, chargerID int64, amount float64, paymentId string, payoutId string) (httpErrorCode int, error error) {
	sqlStatement := `UPDATE charging_processes SET amount=$3, payment_id=$4, payout_id=$5 WHERE (charger_id=$1 AND user_id=(SELECT id FROM users WHERE email=$2) AND amount IS NULL) RETURNING id`
	var id int64
	err := database.DB.QueryRow(sqlStatement, chargerID, userEmail, amount, paymentId, payoutId).Scan(&id)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}
	return http.StatusOK, nil
}

func setChargerAvailable(chargerID int64) (httpErrorCode int, error error) {
	err := charging.UpdateChargerAvailability(chargerID, true)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}
	return http.StatusOK, nil
}

func getCost(chargerID int64, chargedEnergyKWH float64) (cost models.Cost, err error) {
	costPerKWH, err := charging.GetCostPerKWH(chargerID)
	if err != nil {
		return models.Cost{}, fmt.Errorf("internal error")
	}

	return models.Cost{
		Amount:   float32(float64(costPerKWH.Amount) * chargedEnergyKWH),
		Currency: costPerKWH.Currency,
	}, nil
}

func getPayoutAmount(cost models.Cost) (payoutAmount models.Cost) {
	const factor = 0.8

	return models.Cost{
		Amount:   cost.Amount * factor,
		Currency: cost.Currency,
	}
}
