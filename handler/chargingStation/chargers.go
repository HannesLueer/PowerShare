package chargingStation

import (
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func OverviewHandler(w http.ResponseWriter, r *http.Request) {
	chargers, err := getAllCharger()
	if err != nil {
		fmt.Println(err)
	}

	jsonResp, err := json.Marshal(chargers)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(jsonResp)
	return
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	charger, err := getCharger(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResp, err := json.Marshal(charger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
	return
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var charger models.Charger
	err := json.NewDecoder(r.Body).Decode(&charger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var id = CreateCharger(charger)

	jsonResp, err := json.Marshal(id)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(jsonResp)
	return
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var charger models.Charger
	err := json.NewDecoder(r.Body).Decode(&charger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := updateCharger(charger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonResp, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
	return
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	deleteCharger(id)

	return
}
