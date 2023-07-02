package model

type Currency struct {
	Id           string  `json:"id"`
	CurrentPrice float64 `json:"current_price"`
}
