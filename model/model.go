package model

import ("time")

type Currency int

type Operation struct {
	Id			string		`json:"id"`
	Quantity 	float64		`json:"quantity"`
	BuyOrder	Order 		`json:"buy_order"`
	SellOrder	Order 		`json:"sell_order"`
	Currency	Currency 	`json:"currency"`
}

type Order struct {
	TotalPrice	float64		`json:"total_price"`
	UnitPrice	float64		`json:"unit_price"`
	EuroPrice	float64		`json:"euro_price"`
	Date		time.Time	`json:"date"`
}

type Budget struct {
	Currency Currency
	Total 	 float64
}

const (
	Bitcoin Currency = iota
	Euro
	Ether
)