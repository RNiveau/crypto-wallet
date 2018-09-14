package utils

import (
	"github.com/rniveau/crypto-wallet/client"
	"github.com/rniveau/crypto-wallet/model"
)

func GetEuroBudget() *model.Budget {
	euroBudget := client.GetBudgetByCurrency(model.Euro)
	if euroBudget == nil {
		euro := model.Euro
		euroBudget = &model.Budget{Currency: &euro, Total: float64(0)}
	}
	return euroBudget
}

func GetOrderFromOperation(operation *model.Operation) *model.Order {
	if operation.BuyOrder != nil {
		return operation.BuyOrder
	}
	return operation.SellOrder
}