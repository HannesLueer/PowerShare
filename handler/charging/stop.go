package charging

import (
	"PowerShare/database"
	"PowerShare/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func StopHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//stopCharging()

	//json.NewEncoder(w).Encode("")
	return
}

func stopCharging(chargerID int64, userEmail string) (httpErrorCode int, error error) {
	//TODO

	//ladevorgang in db updaten

	// strom ausschalten
	// zähler ablesen
	// kosten errechnen
	// zahlung updaten

	// TODO change isAvailable to false

	return http.StatusOK, nil
}

func writeAmountDB(userEmail string, chargerID int64, amount float32) (httpErrorCode int, error error) {
	sqlStatement := `UPDATE charging_processes SET amount=$3 WHERE (chargerid=$1 AND userid=(SELECT id FROM users WHERE email=$2) AND amount=null)`
	var id int64
	err := database.DB.QueryRow(sqlStatement, chargerID, userEmail, amount).Scan(&id)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}
	return http.StatusOK, nil
}

func captureFundsViaPaypal(paypalOrderID string, price float32, currency string) (err error) {
	// define request
	headerValues := map[string]string{
		"Authorization": " Basic <client_id:secret>",
		"Content-Type":  "application/json",
	}
	paypalUrl := fmt.Sprintf("https://api-m.sandbox.paypal.com/v2/checkout/orders/%s", paypalOrderID)
	httpBodyJson, err := json.Marshal(models.PaypalPatchOrderBody{
		Op:    "replace",
		Path:  "/purchase_units/@reference_id=='default'/amount",
		Value: nil, //TODO set this object
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, paypalUrl, bytes.NewBuffer(httpBodyJson))
	if err != nil {
		return err
	}

	// add header
	for key, value := range headerValues {
		req.Header.Add(key, value)
	}

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// read response
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))

	return nil
}
