package paypal

import (
	"PowerShare/models"
	"encoding/json"
	"fmt"
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
		if accessToken.Ticker != nil {
			accessToken.Ticker.Stop()
			accessToken.Ticker = nil
		}
	}()

	if accessToken.Ticker != nil {
		for true {
			select {
			case <-accessToken.Ticker.C:
				updateAccessToken()
			}
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
