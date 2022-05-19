package shelly

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

var accessToken = models.ShellyAccessToken{
	Token:  "",
	Ticker: nil,
}

func getAccessToken() string {
	if accessToken.Ticker == nil {
		go renewAccessToken()
	}
	return accessToken.Token
}

func getNewAccessToken() (token string, duration time.Duration, err error) {
	// define request
	headerValues := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	urlValues := url.Values{}
	urlValues.Set("itg", os.Getenv("SHELLY_INTEGRATOR_TAG"))
	urlValues.Set("token", os.Getenv("SHELLY_INTEGRATOR_TOKEN"))
	url := fmt.Sprintf("%s/integrator/get_access_token?%s", os.Getenv("SHELLY_API_BASE_URL"), urlValues.Encode())

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(urlValues.Encode()))
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
	respBodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", -1, err
	}

	var respBody models.ShellyAccessTokenResponse
	err = json.Unmarshal(respBodyStr, &respBody)
	if !respBody.Isok {
		err = fmt.Errorf("shelly says there was an error: %q", respBody.Errors)
	}

	return respBody.Data, 24 * time.Hour, err
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
