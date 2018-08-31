package model

import (
	"errors"
	"time"
	"gopkg.in/mgo.v2/bson"

)

type Currency int

type ValidModel interface {
	Valid() error
}

type Operation struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Quantity    float64       `json:"quantity"`
	BuyOrder    *Order        `json:"buy_order"`
	SellOrder   *Order        `json:"sell_order"`
	Currency    *Currency     `json:"currency"`
	Description string        `json:"description"`
}

func (operation *Operation) Valid() error {
	if operation.BuyOrder == nil {
		return errors.New("BuyOrder can't be nil")
	}
	if operation.Quantity <= 0 {
		return errors.New("Quantity must be more than 0")
	}
	if operation.Currency == nil {
		return errors.New("Currency can't be nil")
	}
	if *operation.Currency < Bitcoin || *operation.Currency >= End {
		return errors.New("Currency is not valid")
	}
	return nil
}

type Order struct {
	TotalPrice float64   `json:"total_price"`
	UnitPrice  float64   `json:"unit_price"`
	EuroPrice  float64   `json:"euro_price"`
	Date       time.Time `json:"date"`
}

type Budget struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Currency *Currency     `json:"currency"`
	Total    float64       `json:"total"`
}

const (
	Bitcoin Currency = iota + 1
	Euro
	Ether
	End
)
