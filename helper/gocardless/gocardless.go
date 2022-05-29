package gocardless

import (
	"PowerShare/models"
	"context"
	"encoding/json"
	"fmt"
	gocardless "github.com/gocardless/gocardless-pro-go"
	gocardless2 "github.com/gocardless/gocardless-pro-go/v2"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var client *gocardless.Service
var client2 *gocardless2.Service

func setup() {
	opts := gocardless.WithEndpoint(gocardless.SandboxEndpoint)
	client, _ = gocardless.New(os.Getenv("GOCARDLESS_ACCESS_TOKEN"), opts)
}

func setup2() {
	config, err := gocardless2.NewConfig(os.Getenv("GOCARDLESS_ACCESS_TOKEN"), gocardless2.WithEndpoint(gocardless.SandboxEndpoint))
	if err != nil {
		fmt.Printf("got err in initialising config: %s", err.Error())
		return
	}
	client2, err = gocardless2.New(config)
	if err != nil {
		fmt.Printf("error in initialisating client: %s\n", err.Error())
		return
	}
}

// CreatingPayment creates a new payment for the mandate with the specified id
func CreatingPayment(cost models.Cost, mandateId string) (paymentId string, err error) {
	if client == nil {
		setup()
	}

	ctx := context.TODO()
	paymentCreateParams := gocardless.PaymentCreateParams{}
	paymentCreateParams.Amount = int(cost.Amount * 100) // 1000 --> 10 GBP in pence
	paymentCreateParams.Currency = cost.Currency.Abbreviation
	paymentCreateParams.Links.Mandate = mandateId

	requestOption := gocardless.WithIdempotencyKey(uuid.New().String())
	payment, err := client.Payments.Create(ctx, paymentCreateParams, requestOption)

	return payment.Id, err
}

// CreatingPaymentV2 creates a new payment for the mandate with the specified id using API v2
func CreatingPaymentV2(cost models.Cost, mandateId string) (paymentId string, err error) {
	if client2 == nil {
		setup2()
	}

	paymentCreateParams := gocardless2.PaymentCreateParams{
		Amount:   int(cost.Amount * 100), // 1000 --> 10 GBP in pence
		Currency: cost.Currency.Abbreviation,
		Links: gocardless2.PaymentCreateParamsLinks{
			Mandate: mandateId,
		},
	}

	payment, err := client2.Payments.Create(context.TODO(), paymentCreateParams)

	if err != nil {
		return "", err
	}
	return payment.Id, nil
}

// NewMandate returns the url to a new mandate
func NewMandate(emailAddress string) (url string, err error) {
	brId, err := createBillingRequest()
	if err != nil {
		return "", err
	}

	PresetEmail(brId, emailAddress)

	url, err = createBillingRequestFlow(brId)
	return url, err
}

func createBillingRequest() (billingRequestId string, err error) {
	// define request
	headerValues := map[string]string{
		"GoCardless-Version": "2015-07-06",
		"Authorization":      fmt.Sprintf("Bearer %s", os.Getenv("GOCARDLESS_ACCESS_TOKEN")),
		"Content-Type":       "application/json",
	}
	apiUrl := fmt.Sprintf("%s/billing_requests", os.Getenv("GOCARDLESS_API_BASE_URL"))

	body, err := json.Marshal(
		models.BillingRequestBody{
			BillingRequests: models.BillingRequests{
				MandateRequest: models.MandateRequest{
					Scheme: "sepa_core",
				},
			},
		})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(string(body)))
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
	respBodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respBody models.BillingRequestResponseBody
	err = json.Unmarshal(respBodyStr, &respBody)

	return respBody.BillingRequests.Id, err
}

func createBillingRequestFlow(brId string) (authorisationUrl string, err error) {
	// define request
	headerValues := map[string]string{
		"GoCardless-Version": "2015-07-06",
		"Authorization":      fmt.Sprintf("Bearer %s", os.Getenv("GOCARDLESS_ACCESS_TOKEN")),
		"Content-Type":       "application/json",
	}
	apiUrl := fmt.Sprintf("%s/billing_request_flows", os.Getenv("GOCARDLESS_API_BASE_URL"))

	body, err := json.Marshal(
		models.BillingRequestFlowsRequestBody{
			BillingRequestFlows: models.BillingRequestFlows{
				RedirectUri: os.Getenv("GOCARDLESS_REDIRECT_URL"),
				ExitUri:     os.Getenv("GOCARDLESS_EXIT_URL"),
				Links: models.BillingRequestFlowsLinks{
					BillingRequest: brId,
				},
			},
		})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(string(body)))
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
	respBodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respBody models.BillingRequestFlowsResponseBody
	err = json.Unmarshal(respBodyStr, &respBody)

	return respBody.BillingRequestFlows.AuthorisationUrl, err
}

// GetMandateIdFromEmail returns fist mandate id of customer with given email address
func GetMandateIdFromEmail(emailAddress string) (mandateId string, err error) {
	customerId, err := GetCustomerId(emailAddress)
	if err != nil {
		return "", err
	}

	mandateId, err = GetMandateId(customerId)

	return mandateId, err
}

// GetMandateId returns fist mandate id of customer with given id
func GetMandateId(customerID string) (string, error) {
	if client2 == nil {
		setup2()
	}

	mandateListParams := gocardless2.MandateListParams{
		Customer: customerID,
		Scheme: []string{
			"sepa_core",
		},
	}

	mandateListResult, err := client2.Mandates.List(context.TODO(), mandateListParams)
	if err != nil {
		return "", err
	}

	if len(mandateListResult.Mandates) == 0 {
		return "", fmt.Errorf("no mandate was found with the scheme sepa")
	}

	return mandateListResult.Mandates[0].Id, nil
}

// GetCustomerId returns fist customer id of customer with given email
func GetCustomerId(emailAddress string) (string, error) {
	if client2 == nil {
		setup2()
	}

	customerListParams := gocardless2.CustomerListParams{}

	customerListResult, err := client2.Customers.List(context.TODO(), customerListParams)
	if err != nil {
		return "", err
	}

	for _, customer := range customerListResult.Customers {
		if customer.Email == emailAddress {
			return customer.Id, nil
		}
	}

	return "", fmt.Errorf("no customer found")
}

// PresetEmail already sets the email address of a billing request
func PresetEmail(billingRequestId string, emailAddress string) {
	if client2 == nil {
		setup2()
	}

	billingRequestCollectCustomerDetailsParams := gocardless2.BillingRequestCollectCustomerDetailsParams{
		Customer: &gocardless2.BillingRequestCollectCustomerDetailsParamsCustomer{
			Email: emailAddress,
		},
	}

	client2.BillingRequests.CollectCustomerDetails(context.TODO(), billingRequestId, billingRequestCollectCustomerDetailsParams)
}
