package paypal

import (
	"PowerShare/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var accessToken = models.PaypalAccessToken{
	Token:  "",
	Ticker: nil,
}

func CreatePaypalOrder() (respBody string, err error) {
	// define request
	headerValues := map[string]string{
		"Authorization":     fmt.Sprintf("Bearer %s", GetAccessToken()),
		"PayPal-Request-Id": uuid.New().String(),
		"Content-Type":      "application/json",
	}
	paypalUrl := fmt.Sprintf("%s/v2/checkout/orders/", os.Getenv("PAYPAL_API_BASE_URL"))
	httpBodyJson, err := json.Marshal(models.PaypalNewOrder{
		Intent: "AUTHORIZE",
		PurchaseUnits: []models.PaypalPurchaseUnit{
			{
				Amount: models.PaypalAmountValueStr{
					CurrencyCode: "EUR",
					Value:        "10.0",
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	//debug
	httpBodyJson = []byte("{\n    \"intent\": \"AUTHORIZE\",\n    \"purchase_units\": [\n        {\n            \"items\": [\n                {\n                    \"name\": \"T-Shirt\",\n                    \"description\": \"Green XL\",\n                    \"quantity\": \"1\",\n                    \"unit_amount\": {\n                        \"currency_code\": \"USD\",\n                        \"value\": \"100.00\"\n                    }\n                }\n            ],\n            \"amount\": {\n                \"currency_code\": \"USD\",\n                \"value\": \"100.00\",\n                \"breakdown\": {\n                    \"item_total\": {\n                        \"currency_code\": \"USD\",\n                        \"value\": \"100.00\"\n                    }\n                }\n            }\n        }\n    ],\n    \"application_context\": {\n        \"return_url\": \"https://example.com/return\",\n        \"cancel_url\": \"https://example.com/cancel\"\n    }\n}")

	req, err := http.NewRequest(http.MethodPost, paypalUrl, bytes.NewBuffer(httpBodyJson))
	if err != nil {
		return "", err
	}

	// add header
	for key, value := range headerValues {
		req.Header.Add(key, value)
	}

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// read response
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

func UpdateOrderPaypal(paypalOrderID string, cost models.Cost) (err error) {
	// define request
	headerValues := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", GetAccessToken()),
		"Content-Type":  "application/json",
	}
	paypalUrl := fmt.Sprintf("%s/v2/checkout/orders/%s", os.Getenv("PAYPAL_API_BASE_URL"), paypalOrderID)

	httpBodyJson, err := json.Marshal([]models.PaypalPatchOrderBody{
		{
			Op:   "replace",
			Path: "/purchase_units/@reference_id=='default'/amount",
			Value: models.PaypalAmountValue{
				CurrencyCode: cost.Currency.Abbreviation,
				Value:        cost.Amount,
			},
		},
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

	// TODO: fix error (order "status":"COMPLETED")
	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))

	return nil
}

func CaptureFundsPaypal(paypalOrderID string) (err error) {
	// define request
	headerValues := map[string]string{
		"Authorization":     fmt.Sprintf("Bearer %s", GetAccessToken()),
		"Content-Type":      "application/json",
		"PayPal-Request-Id": uuid.New().String(),
	}
	paypalUrl := fmt.Sprintf("%s/v2/payments/authorizations/%s/capture ", os.Getenv("PAYPAL_API_BASE_URL"), paypalOrderID)
	req, err := http.NewRequest(http.MethodPost, paypalUrl, nil)
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
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("paypal: %s", resp.Status)
	}

	return nil
}

func GetPaypalOrder(paypalOrderID string) (responseBody string, paypalStatusCode int, err error) {
	// define request
	headerValues := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", GetAccessToken()),
		"Content-Type":  "application/json",
	}
	paypalUrl := fmt.Sprintf("%s/v2/checkout/orders/%s", os.Getenv("PAYPAL_API_BASE_URL"), paypalOrderID)

	req, err := http.NewRequest(http.MethodGet, paypalUrl, nil)
	if err != nil {
		return "", -1, err
	}

	// add header
	for key, value := range headerValues {
		req.Header.Add(key, value)
	}

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", -1, err
	}

	// read response
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", -1, err
	}

	return string(respBody), resp.StatusCode, nil
}

func IsOrderValid(paypalOrderID string) bool {
	_, paypalStatusCode, err := GetPaypalOrder(paypalOrderID)
	return err == nil && paypalStatusCode == http.StatusOK
}

func GetAccessToken() string {
	if accessToken.Ticker == nil {
		renewAccessToken()
	}
	return accessToken.Token
}

func getNewAccessToken() (token string, duration time.Duration, err error) {
	// define request
	headerValues := map[string]string{
		"Accept":          "application/json",
		"Accept-Language": "en_US",
		"Content-Type":    "application/x-www-form-urlencoded",
	}
	urlValues := url.Values{}
	urlValues.Set("grant_type", "client_credentials")
	paypalUrl := fmt.Sprintf("%s/v1/oauth2/token?%s", os.Getenv("PAYPAL_API_BASE_URL"), urlValues.Encode())

	req, err := http.NewRequest(http.MethodPost, paypalUrl, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return "", -1, err
	}

	// add header
	for key, value := range headerValues {
		req.Header.Add(key, value)
	}

	// add auth
	req.SetBasicAuth(os.Getenv("PAYPAL_CLIENT_ID"), os.Getenv("PAYPAL_SECRET"))

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", -1, err
	}

	// read response
	defer resp.Body.Close()
	respBodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", -1, err
	}

	var respBody models.PaypalAccessTokenResponse
	err = json.Unmarshal(respBodyStr, &respBody)

	return respBody.AccessToken, time.Duration(respBody.ExpiresIn) * time.Second, err
}

func renewAccessToken() {
	updateAccessToken()

	defer func() {
		accessToken.Ticker.Stop()
		accessToken.Ticker = nil
	}()

	for true {
		select {
		case <-accessToken.Ticker.C:
			updateAccessToken()
		}
	}
}

func updateAccessToken() {
	const loadAhead = time.Minute

	token, duration, err := getNewAccessToken()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	accessToken.Token = token
	accessToken.Ticker = time.NewTicker(duration - loadAhead)
}
