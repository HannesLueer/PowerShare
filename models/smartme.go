package models

import "time"

type SmartmeAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

type SmartmeDevice struct {
	Id                          string    `json:"Id"`
	Name                        string    `json:"Name"`
	Serial                      int       `json:"Serial"`
	DeviceEnergyType            string    `json:"DeviceEnergyType"`
	MeterSubType                string    `json:"MeterSubType"`
	FamilyType                  string    `json:"FamilyType"`
	ActivePower                 float64   `json:"ActivePower"`
	ActivePowerL1               float64   `json:"ActivePowerL1"`
	ActivePowerL2               float64   `json:"ActivePowerL2"`
	ActivePowerL3               float64   `json:"ActivePowerL3"`
	ActivePowerUnit             string    `json:"ActivePowerUnit"`
	CounterReading              float64   `json:"CounterReading"`
	CounterReadingUnit          string    `json:"CounterReadingUnit"`
	CounterReadingT1            float64   `json:"CounterReadingT1"`
	CounterReadingT2            float64   `json:"CounterReadingT2"`
	CounterReadingT3            float64   `json:"CounterReadingT3"`
	CounterReadingT4            float64   `json:"CounterReadingT4"`
	CounterReadingImport        float64   `json:"CounterReadingImport"`
	CounterReadingExport        float64   `json:"CounterReadingExport"`
	SwitchOn                    bool      `json:"SwitchOn"`
	SwitchPhaseL1On             bool      `json:"SwitchPhaseL1On"`
	SwitchPhaseL2On             bool      `json:"SwitchPhaseL2On"`
	SwitchPhaseL3On             bool      `json:"SwitchPhaseL3On"`
	Voltage                     float64   `json:"Voltage"`
	VoltageL1                   float64   `json:"VoltageL1"`
	VoltageL2                   float64   `json:"VoltageL2"`
	VoltageL3                   float64   `json:"VoltageL3"`
	Current                     float64   `json:"Current"`
	CurrentL1                   float64   `json:"CurrentL1"`
	CurrentL2                   float64   `json:"CurrentL2"`
	CurrentL3                   float64   `json:"CurrentL3"`
	PowerFactor                 float64   `json:"PowerFactor"`
	PowerFactorL1               float64   `json:"PowerFactorL1"`
	PowerFactorL2               float64   `json:"PowerFactorL2"`
	PowerFactorL3               float64   `json:"PowerFactorL3"`
	Temperature                 float64   `json:"Temperature"`
	ActiveTariff                int       `json:"ActiveTariff"`
	DigitalOutput1              bool      `json:"DigitalOutput1"`
	DigitalOutput2              bool      `json:"DigitalOutput2"`
	AnalogOutput1               int       `json:"AnalogOutput1"`
	AnalogOutput2               int       `json:"AnalogOutput2"`
	DigitalInput1               bool      `json:"DigitalInput1"`
	DigitalInput2               bool      `json:"DigitalInput2"`
	ValueDate                   time.Time `json:"ValueDate"`
	AdditionalMeterSerialNumber string    `json:"AdditionalMeterSerialNumber"`
	FlowRate                    float64   `json:"FlowRate"`
	ChargingStationState        string    `json:"ChargingStationState"`
}
