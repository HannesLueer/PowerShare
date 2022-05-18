package models

import "time"

type PaypalPatchOrderBody struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type PaypalAmountValue struct {
	CurrencyCode string  `json:"currency_code"`
	Value        float32 `json:"value"`
}

type PaypalAmountValueStr struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type PaypalAccessToken struct {
	Token  string
	Ticker *time.Ticker
}

type PaypalAccessTokenResponse struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	AppId       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

type PaypalNewOrder struct {
	Intent        string               `json:"intent"`
	PurchaseUnits []PaypalPurchaseUnit `json:"purchase_units"`
	PaymentSource PaypalPaymentSource  `json:"payment_source"`
}

type PaypalPurchaseUnit struct {
	ReferenceId string               `json:"reference_id"`
	Amount      PaypalAmountValueStr `json:"amount"`
	Payee       struct {
		EmailAddress string `json:"email_address"`
	} `json:"payee"`
	PaymentInstruction struct {
		PlatformFees []struct {
			Amount struct {
				CurrencyCode string `json:"currency_code"`
				Value        string `json:"value"`
			} `json:"amount"`
			Payee struct {
				EmailAddress string `json:"email_address"`
			} `json:"payee"`
		} `json:"platform_fees"`
		DisbursementMode   string `json:"disbursement_mode"`
		PayeePricingTierId string `json:"payee_pricing_tier_id"`
	} `json:"payment_instruction"`
}

type PaypalPaymentSource struct {
	ApplePay struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		EmailAddress string `json:"email_address"`
		PhoneNumber  struct {
			CountryCode    string `json:"country_code"`
			NationalNumber string `json:"national_number"`
		} `json:"phone_number"`
		Shipping struct {
			Name struct {
				GivenName string `json:"given_name"`
				Surname   string `json:"surname"`
			} `json:"name"`
			EmailAddress string `json:"email_address"`
			Address      struct {
				AddressLine1 string `json:"address_line_1"`
				AddressLine2 string `json:"address_line_2"`
				AdminArea2   string `json:"admin_area_2"`
				AdminArea1   string `json:"admin_area_1"`
				PostalCode   string `json:"postal_code"`
				CountryCode  string `json:"country_code"`
			} `json:"address"`
		} `json:"shipping"`
		DecryptedToken struct {
			TransactionAmount struct {
				CurrencyCode string `json:"currency_code"`
				Value        string `json:"value"`
			} `json:"transaction_amount"`
			TokenizedCard struct {
				Number         string `json:"number"`
				Expiry         string `json:"expiry"`
				BillingAddress struct {
					AddressLine1 string `json:"address_line_1"`
					AddressLine2 string `json:"address_line_2"`
					AdminArea2   string `json:"admin_area_2"`
					AdminArea1   string `json:"admin_area_1"`
					PostalCode   string `json:"postal_code"`
					CountryCode  string `json:"country_code"`
				} `json:"billing_address"`
			} `json:"tokenized_card"`
			DeviceManufacturerId string `json:"device_manufacturer_id"`
			PaymentDataType      string `json:"payment_data_type"`
			PaymentData          struct {
				Cryptogram   string `json:"cryptogram"`
				EciIndicator string `json:"eci_indicator"`
			} `json:"payment_data"`
		} `json:"decrypted_token"`
	} `json:"apple_pay"`
}
