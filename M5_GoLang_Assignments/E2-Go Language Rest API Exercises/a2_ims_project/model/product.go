package model

type Product struct {
	ID       int     `json:"id"`
	Name     int     `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Category string  `json:"category"`
}
