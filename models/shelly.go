package models

import "time"

type ShellyAccessToken struct {
	Token  string
	Ticker *time.Ticker
}

type ShellyAccessTokenResponse struct {
	Isok   bool     `json:"isok"`
	Data   string   `json:"data"`
	Errors []string `json:"errors"`
}

type ShellyOnOffCommandRequest struct {
	Event    string                        `json:"event"`    // EVENT_NAME
	TrId     int                           `json:"trid"`     // TRANSACTION_ID
	DeviceId int                           `json:"deviceId"` // DEVICE_ID
	Data     ShellyOnOffCommandRequestData `json:"data"`
}

type ShellyOnOffCommandRequestData struct {
	Cmd    string                              `json:"cmd"` // COMMAND_NAME "relay" | "light"
	Params ShellyOnOffCommandRequestDataParams `json:"params"`
}

type ShellyOnOffCommandRequestDataParams struct {
	Id   int    `json:"id"`   // DEVICE_CHANNEL
	Turn string `json:"turn"` // COMMAND_ACTION "on" | "off"
}
