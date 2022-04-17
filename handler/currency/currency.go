package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func OverviewHandler(w http.ResponseWriter, r *http.Request) {
	currencies, err := getCurrencies()
	if err != nil {
		fmt.Println(err)
	}

	jsonResp, err := json.Marshal(currencies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
	return
}
