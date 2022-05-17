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

	//debug
	fmt.Println("11111111111111")
	fmt.Println("GetPaypalOrder: ")
	resp, _, err := helper.GetPaypalOrder(paypalOrderID)
	fmt.Println(resp)
	fmt.Println(err)
	fmt.Println("--------------")

	helper.UpdateOrderPaypal(paypalOrderID, models.Cost{
		Amount: 123.45,
		Currency: models.Currency{
			Abbreviation: "USD",
			Symbol:       "$",
		},
	})
	fmt.Println("--------------")

	fmt.Println("22222222222222")
	resp, _, _ = helper.GetPaypalOrder(paypalOrderID)
	fmt.Println(resp)
	//debug end

	return
}

func NewOrderHandler(w http.ResponseWriter, r *http.Request) {
	respStr, err := helper.CreatePaypalOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(respStr))
}

func startCharging(chargerID int64, userEmail string, paypalOrderID string) (httpErrorCode int, error error) {
	// check if charging station is available
	isChargerAvailable, err := helper.IsChargerAvailable(chargerID)
	if err != nil || !isChargerAvailable {
		return http.StatusBadRequest, fmt.Errorf("charger is not available")
	}

	// TODO activate / read electricity meter

	// turn power on
	err = helper.SwitchPower(chargerID, true)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// write loading process in db
	writeChargingProcessDB(userEmail, chargerID, paypalOrderID)

	return http.StatusOK, nil
}

func writeChargingProcessDB(userEmail string, chargerID int64, paypalOrderID string) (httpErrorCode int, error error) {
	// add charging process
	sqlStatement := `INSERT INTO charging_processes (userid, chargerid, paypal_order_id) VALUES (((SELECT id FROM users WHERE email=$1)), $2, $3) RETURNING id`
	var id int64
	err := database.DB.QueryRow(sqlStatement, userEmail, chargerID, paypalOrderID).Scan(&id)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}

	// set charger occupied
	err = helper.UpdateChargerAvailability(chargerID, false)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}

	return http.StatusOK, nil
}
