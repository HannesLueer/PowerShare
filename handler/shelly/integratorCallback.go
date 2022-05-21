package shelly

import (
	"PowerShare/helper/shelly"
	"PowerShare/models"
	"encoding/json"
	"fmt"
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
		return shelly.AddDeviceToDB(body)
	}

	if body.Action == "remove" {
		return shelly.RemoveDeviceFromDB(body)
	}

	return http.StatusBadRequest, fmt.Errorf("action must have value \"add\" or \"remove\"")
}
