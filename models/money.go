package models

type Currency struct {
	ID           int64  `json:"-"`
	Abbreviation string `json:"abbreviation"`
	Symbol       string `json:"symbol"`
}

type Cost struct {
	Amount   float32  `json:"amount"`
	Currency Currency `json:"currency"`
}
