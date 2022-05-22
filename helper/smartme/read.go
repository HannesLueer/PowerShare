package smartme

import (
	"PowerShare/helper/charger"
	"PowerShare/helper/user"
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func ReadCounter(userEmail string, chargerID int64) (counterKWH float64, err error) {
	id, err := user.GetId(userEmail)
	if err != nil {
		return -1, err
	}

	serialNumber := charger.GetSmartmeSerialNumber(chargerID)

	return readCounter(id, serialNumber)
}

// https://smart-me.com/swagger/ui/index#!/DeviceBySerial/DeviceBySerial_Get
func readCounter(userId int64, serialNumber string) (counterKWH float64, err error) {
	// get access token
	token, err := GetAccessToken(userId)

	// define request
	headerValues := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
	}
	urlValues := url.Values{}
	urlValues.Set("serial", serialNumber)
	apiUrl := fmt.Sprintf("%s/DeviceBySerial%s", os.Getenv("SMARTME_API_BASE_URL"), urlValues.Encode())

	req, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return -1, err
	}

	// add header
	for key, value := range headerValues {
		req.Header.Add(key, value)
	}

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode == http.StatusUnauthorized {
		// try again with new token if unauthorized
		GetAndSaveNewAccessToken(userId)
		return readCounter(userId, serialNumber)
	}
	if err != nil {
		return -1, err
	}

	// read response
	defer resp.Body.Close()
	respBodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	var device models.SmartmeDevice
	err = json.Unmarshal(respBodyStr, &device)

	// calculate kwh
	counterKWH, err = calculateAmountInKWH(device.CounterReading, device.CounterReadingUnit)
	if err != nil {
		return -1, err
	}

	return counterKWH, err
}

func calculateAmountInKWH(value float64, unit string) (kwh float64, err error) {
	unit = strings.ToLower(unit)
	var factor = 1.0
	switch unit {
	case "wh":
		factor = 0.001
	case "kwh":
		factor = 1.0
	case "mwh":
		factor = 1_000
	case "gwh":
		factor = 1_000_000
	default:
		err = fmt.Errorf("unknown unit")
	}

	return value * factor, err
}
