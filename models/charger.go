package models

type Charger struct {
	ID            int64         `json:"id"`
	Title         string        `json:"title"`
	Position      Coordinate    `json:"position"`
	Cost          Cost          `json:"cost"`
	IsOccupied    bool          `json:"isOccupied"`
	Description   string        `json:"description"`
	TechnicalData TechnicalData `json:"technicalData"`
}

type Coordinate struct {
	Lat float64
	Lng float64
}

type TechnicalData struct {
	ShellyDeviceId int64 `json:"shellyDeviceId"`
}
