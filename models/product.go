package models

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	BasePrice  float64 `json:"base_price"`
	FinalPrice float64 `json:"final_price"`
}