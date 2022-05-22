package smartme

import (
	"PowerShare/database"
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func RequestAndSaveFirstAccessToken(userEmail string, authCode string) (err error) {
	accessToken, err := requestFirstAccessToken(authCode)
	if err != nil {
		return err
	}

	err = writeNewAccessTokenToDB(userEmail, accessToken)
	return err
}

func requestFirstAccessToken(authCode string) (accessToken models.SmartmeAccessToken, err error) {
	// define request
	headerValues := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	urlValues := url.Values{}
	urlValues.Set("client_id", os.Getenv("SMARTME_CLIENT_ID"))
	urlValues.Set("client_secret", os.Getenv("SMARTME_CLIENT_SECRET"))
	urlValues.Set("grant_type", "authorization_code")
	urlValues.Set("code", authCode)
	urlValues.Set("redirect_uri", "")
	apiUrl := fmt.Sprintf("%s/oauth/token/%s", os.Getenv("SMARTME_API_BASE_URL"), urlValues.Encode())

	req, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return accessToken, err
	}

	// add header
	for key, value := range headerValues {
		req.Header.Add(key, value)
	}

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return accessToken, err
	}

	// read response
	defer resp.Body.Close()
	respBodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return accessToken, err
	}

	err = json.Unmarshal(respBodyStr, &accessToken)

	return accessToken, err
}

func GetAndSaveNewAccessToken(userId int64) (accessToken models.SmartmeAccessToken, err error) {
	accessToken, err = getNewAccessToken(userId)
	if err != nil {
		return accessToken, err
	}

	err = updateAccessTokenInDB(userId, accessToken)
	return accessToken, err
}

func getNewAccessToken(userId int64) (accessToken models.SmartmeAccessToken, err error) {
	// get refresh token
	oldAccessToken, err := GetAccessToken(userId)
	if err != nil {
		return accessToken, err
	}

	// define request
	headerValues := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	urlValues := url.Values{}
	urlValues.Set("client_id", os.Getenv("SMARTME_CLIENT_ID"))
	urlValues.Set("client_secret", os.Getenv("SMARTME_CLIENT_SECRET"))
	urlValues.Set("response_type", "refresh_token")
	urlValues.Set("refresh_token", oldAccessToken.RefreshToken)
	apiUrl := fmt.Sprintf("%s/oauth/token/%s", os.Getenv("SMARTME_API_BASE_URL"), urlValues.Encode())

	req, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return accessToken, err
	}

	// add header
	for key, value := range headerValues {
		req.Header.Add(key, value)
	}

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return accessToken, err
	}

	// read response
	defer resp.Body.Close()
	respBodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return accessToken, err
	}

	err = json.Unmarshal(respBodyStr, &accessToken)

	return accessToken, err
}

func GetAccessToken(userId int64) (token models.SmartmeAccessToken, err error) {
	sqlStatement := `SELECT access_token, token_type, refresh_token FROM smartme_tokens WHERE user_id=$1`
	err = database.DB.QueryRow(sqlStatement, userId).Scan(&token.AccessToken, &token.TokenType, &token.RefreshToken)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return token, fmt.Errorf("internal error")
	}
	return token, nil
}

func writeNewAccessTokenToDB(userEmail string, accessToken models.SmartmeAccessToken) (err error) {
	var userId int64
	sqlInsertStatement := `INSERT INTO smartme_tokens (user_id, access_token, token_type, refresh_token) VALUES ((SELECT id FROM users WHERE email=$1), $2, $3, $4) RETURNING user_id`
	err = database.DB.QueryRow(sqlInsertStatement, userEmail, accessToken.AccessToken, accessToken.TokenType, accessToken.RefreshToken).Scan(&userId)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return fmt.Errorf("internal error")
	}
	return nil
}

func updateAccessTokenInDB(userId int64, accessToken models.SmartmeAccessToken) (err error) {
	sqlInsertStatement := `UPDATE smartme_tokens SET access_token=$2, token_type=$3, refresh_token=$4 WHERE user_id=$1 RETURNING user_id`
	err = database.DB.QueryRow(sqlInsertStatement, userId, accessToken.AccessToken, accessToken.TokenType, accessToken.RefreshToken).Scan(&userId)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return fmt.Errorf("internal error")
	}
	return nil
}
