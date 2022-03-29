package models

type Charger struct {
	ID         int64      `json:"id"`
	Title      string     `json:"title"`
	Position   Coordinate `json:"position"`
	Cost       float32    `json:"cost"`
	IsOccupied bool       `json:"isOccupied"`
}

type Coordinate struct {
	Lat float64
	Lng float64
}
