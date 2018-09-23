package model

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"log"
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
	ParentId	string		  `json:"parent_id" bson:",omitempty"`
	Parent 		*Operation 	  `json:"parent" bson:"-"`
}

func (operation *Operation) Valid() error {
	if operation.Quantity <= 0 {
		return errors.New("Quantity must be more than 0")
	}
	if operation.Currency == nil {
		return errors.New("Currency can't be nil")
	}
	if *operation.Currency < Bitcoin || *operation.Currency >= End {
		return errors.New("Currency is not valid")
	}
	if operation.ParentId != "" {
           log.Println(operation.Parent)
           if operation.Parent == nil {
           	return errors.New("ParentId doesn't exist")
           }
    }
    if operation.BuyOrder == nil && operation.SellOrder == nil {
        return errors.New("You need an order in an operation")
    }
    if operation.BuyOrder != nil {
    	if err := operation.BuyOrder.Valid(); err != nil {
    		return err
    	}
    }
    if operation.SellOrder != nil {
    	if err := operation.SellOrder.Valid(); err != nil {
    		return err
    	}
    }
    return nil
}

type Order struct {
	Price     float64    `json:"price"`
	EuroPrice float64    `json:"euro_price"`
	Currency  *Currency  `json:"currency"`
	Date      customTime `json:"date"`
}

func (order *Order) Valid() error {
	if order.Currency == nil {
		return errors.New("Currency can't be nil")
	}
	if *order.Currency < Bitcoin || *order.Currency >= End {
		return errors.New("Currency is not valid")
	}
	if *order.Currency != Euro && order.EuroPrice <= 0 {
		return errors.New("Euro price must be filled")
	}
	return nil
}

type Budget struct {
	Id       	bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Currency 	*Currency     `json:"currency"`
	Total    	float64       `json:"total"`
	Available   float64       `json:"available"`
}

const (
	Bitcoin Currency = iota + 1
	Euro
	Ether
	Ripple
	IOST
	End
)
