package chargingStation

import (
	chargerHelper "PowerShare/helper/charger"
	"PowerShare/helper/jwt"
	"PowerShare/helper/user"
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
	return
}

func OverviewOwnHandler(w http.ResponseWriter, r *http.Request) {
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

	chargers, err := getChargersOfUser(email)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(chargers)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
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

	// get user email
	tokenStr, _, _ := jwt.GetToken(r)
	email, _ := jwt.GetEmailFromToken(tokenStr)

	// get user id
	userId, err := user.GetId(email)
	if err != nil {
		userId = -1
	}

	charger, err := getCharger(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// include the technical data only if the charger belongs to the user
	if userId == -1 || userId != chargerHelper.GetUserId(charger.ID) {
		charger.TechnicalData = models.TechnicalData{}
	}

	jsonResp, err := json.Marshal(charger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
	return
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
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

	var charger models.Charger
	err = json.NewDecoder(r.Body).Decode(&charger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var id = CreateCharger(charger, email)

	jsonResp, err := json.Marshal(id)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(jsonResp)
	return
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	var charger models.Charger
	err = json.NewDecoder(r.Body).Decode(&charger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := updateCharger(charger, email)
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

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deleteCharger(id, email)

	return
}
