package models

import "time"

type BillingRequestBody struct {
	BillingRequests BillingRequests `json:"billing_requests"`
}

type BillingRequests struct {
	MandateRequest `json:"mandate_request"`
}

type MandateRequest struct {
	Scheme string `json:"scheme"`
}

type BillingRequestResponseBody struct {
	BillingRequests struct {
		Id             string    `json:"id"`
		CreatedAt      time.Time `json:"created_at"`
		Status         string    `json:"status"`
		MandateRequest struct {
			Currency string `json:"currency"`
			Scheme   string `json:"scheme"`
			Verify   string `json:"verify"`
			Links    struct {
			} `json:"links"`
			Metadata struct {
			} `json:"metadata"`
		} `json:"mandate_request"`
		PaymentRequest interface{} `json:"payment_request"`
		Metadata       interface{} `json:"metadata"`
		Links          struct {
			Customer              string `json:"customer"`
			CustomerBillingDetail string `json:"customer_billing_detail"`
			Creditor              string `json:"creditor"`
			MandateRequest        string `json:"mandate_request"`
		} `json:"links"`
		FallbackEnabled  bool   `json:"fallback_enabled"`
		FallbackOccurred bool   `json:"fallback_occurred"`
		CreditorName     string `json:"creditor_name"`
		Actions          []struct {
			Type                   string   `json:"type"`
			Required               bool     `json:"required"`
			CompletesActions       []string `json:"completes_actions"`
			RequiresActions        []string `json:"requires_actions"`
			Status                 string   `json:"status"`
			AvailableCurrencies    []string `json:"available_currencies,omitempty"`
			CollectCustomerDetails struct {
				IncompleteFields struct {
					Customer              []string `json:"customer"`
					CustomerBillingDetail []string `json:"customer_billing_detail"`
				} `json:"incomplete_fields"`
				DefaultCountryCode string `json:"default_country_code"`
			} `json:"collect_customer_details,omitempty"`
			AvailableCountryCodes []string `json:"available_country_codes,omitempty"`
			BankAuthorisation     struct {
				AuthorisationType   string `json:"authorisation_type"`
				RequiresInstitution bool   `json:"requires_institution"`
				DirectToInstitution bool   `json:"direct_to_institution"`
				Adapter             string `json:"adapter"`
			} `json:"bank_authorisation,omitempty"`
		} `json:"actions"`
		Resources struct {
			Customer struct {
				Id          string      `json:"id"`
				CreatedAt   time.Time   `json:"created_at"`
				Email       interface{} `json:"email"`
				GivenName   interface{} `json:"given_name"`
				FamilyName  interface{} `json:"family_name"`
				CompanyName interface{} `json:"company_name"`
				Language    string      `json:"language"`
				PhoneNumber interface{} `json:"phone_number"`
				Metadata    struct {
				} `json:"metadata"`
			} `json:"customer"`
			CustomerBillingDetail struct {
				Id                    string      `json:"id"`
				CreatedAt             time.Time   `json:"created_at"`
				AddressLine1          interface{} `json:"address_line1"`
				AddressLine2          interface{} `json:"address_line2"`
				AddressLine3          interface{} `json:"address_line3"`
				City                  interface{} `json:"city"`
				Region                interface{} `json:"region"`
				PostalCode            interface{} `json:"postal_code"`
				CountryCode           interface{} `json:"country_code"`
				SwedishIdentityNumber interface{} `json:"swedish_identity_number"`
				DanishIdentityNumber  interface{} `json:"danish_identity_number"`
			} `json:"customer_billing_detail"`
		} `json:"resources"`
	} `json:"billing_requests"`
}

type BillingRequestFlowsRequestBody struct {
	BillingRequestFlows BillingRequestFlows `json:"billing_request_flows"`
}

type BillingRequestFlows struct {
	RedirectUri string                   `json:"redirect_uri"`
	ExitUri     string                   `json:"exit_uri"`
	Links       BillingRequestFlowsLinks `json:"links"`
}

type BillingRequestFlowsLinks struct {
	BillingRequest string `json:"billing_request"`
}

type BillingRequestFlowsResponseBody struct {
	BillingRequestFlows struct {
		AuthorisationUrl    string    `json:"authorisation_url"`
		LockCustomerDetails bool      `json:"lock_customer_details"`
		LockBankAccount     bool      `json:"lock_bank_account"`
		AutoFulfil          bool      `json:"auto_fulfil"`
		CreatedAt           time.Time `json:"created_at"`
		ExpiresAt           time.Time `json:"expires_at"`
		RedirectUri         string    `json:"redirect_uri"`
		Links               struct {
			BillingRequest string `json:"billing_request"`
		} `json:"links"`
	} `json:"billing_request_flows"`
}
