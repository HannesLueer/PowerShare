package paypal

import (
	"PowerShare/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func CreatePayoutPaypal(receiverEmail string, cost models.Cost) (respBody models.PayPalPayoutResponse, err error) {
	// define request
	headerValues := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", GetAccessToken()),
		"Content-Type":  "application/json",
	}
	paypalUrl := fmt.Sprintf("%s/v1/payments/payouts", os.Getenv("PAYPAL_API_BASE_URL"))

	httpBodyJson, err := json.Marshal(models.PaypalPayoutBody{
		SenderBatchHeader: models.PaypalPayoutSenderBatchHeader{
			SenderBatchId: uuid.New().String(),
			EmailSubject:  "You have a payout!",
			EmailMessage:  "You have received a payout!",
		},
		Items: []models.PaypalPayoutItem{
			{
				RecipientType: "EMAIL",
				Amount: models.PaypalAmount{
					Value:    strconv.FormatFloat(float64(cost.Amount), 'f', 2, 64),
					Currency: cost.Currency.Abbreviation,
				},
				Note:     "Thanks for making PowerShare great!",
				Receiver: receiverEmail,
			},
		},
	})
	if err != nil {
		return respBody, err
	}

	req, err := http.NewRequest(http.MethodPost, paypalUrl, bytes.NewBuffer(httpBodyJson))
	if err != nil {
		return respBody, err
	}

	// add header
	for key, value := range headerValues {
		req.Header.Add(key, value)
	}

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return respBody, err
	}

	// read response
	defer resp.Body.Close()
	respBodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respBody, err
	}

	err = json.Unmarshal(respBodyStr, &respBody)

	return respBody, err
}
