package models

type Location struct {
	ID 		int64   `json:"id"`
	Country string  `json:"country"`
	City 	string  `json:"city,omitempty"`
	TaxRate float64 `json:"tax_rate"`
	Currency string `json:"currency"`
	Discount float64 `json:"discount,omitempty"`
}