package models

import "time"

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

type PaypalPayoutBody struct {
	SenderBatchHeader PaypalPayoutSenderBatchHeader `json:"sender_batch_header"`
	Items             []PaypalPayoutItem            `json:"items"`
}

type PaypalPayoutSenderBatchHeader struct {
	SenderBatchId string `json:"sender_batch_id"`
	EmailSubject  string `json:"email_subject"`
	EmailMessage  string `json:"email_message"`
}

type PaypalPayoutItem struct {
	RecipientType        string       `json:"recipient_type"`
	Amount               PaypalAmount `json:"amount"`
	Note                 string       `json:"note"`
	SenderItemId         string       `json:"sender_item_id"`
	Receiver             string       `json:"receiver"`
	NotificationLanguage string       `json:"notification_language,omitempty"`
}

type PaypalAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type PayPalPayoutResponse struct {
	BatchHeader struct {
		PayoutBatchId     string `json:"payout_batch_id"`
		BatchStatus       string `json:"batch_status"`
		SenderBatchHeader struct {
			SenderBatchId string `json:"sender_batch_id"`
			EmailSubject  string `json:"email_subject"`
			EmailMessage  string `json:"email_message"`
		} `json:"sender_batch_header"`
	} `json:"batch_header"`
	Links []struct {
		Href    string `json:"href"`
		Rel     string `json:"rel"`
		Method  string `json:"method"`
		EncType string `json:"encType"`
	} `json:"links"`
}
