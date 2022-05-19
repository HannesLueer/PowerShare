package shelly

import (
	"PowerShare/database"
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func IntegratorAddRemoveCallbackHandler(w http.ResponseWriter, r *http.Request) {
	var body models.ShellyAddRemoveDeviceBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "error in reading body", http.StatusBadRequest)
		return
	}

	errCode, err := writeToDB(body)
	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}
	return
}

func writeToDB(body models.ShellyAddRemoveDeviceBody) (errCode int, err error) {
	if body.Action == "add" {
		return addDeviceToDB(body)
	}

	if body.Action == "remove" {
		return removeDeviceFromDB(body)
	}

	return http.StatusBadRequest, fmt.Errorf("action must have value \"add\" or \"remove\"")
}

func addDeviceToDB(body models.ShellyAddRemoveDeviceBody) (errCode int, err error) {
	sqlStatement := `INSERT INTO shelly_devices (device_id, user_id, device_type, device_code, host) VALUES ($1, $2, $3, $4, $5) RETURNING device_id`
	
	_, err = database.DB.Exec(sqlStatement, body.DeviceId, body.UserId, body.DeviceType, body.DeviceCode, body.Host)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("an error has occurred")
	}

	return http.StatusOK, nil
}

func removeDeviceFromDB(body models.ShellyAddRemoveDeviceBody) (errCode int, err error) {
	sqlStatement := `DELETE FROM shelly_devices WHERE device_id = $1 AND user_id=$2`

	_, err = database.DB.Exec(sqlStatement, body.DeviceId, body.UserId)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("an error has occurred")
	}

	return http.StatusOK, nil
}
