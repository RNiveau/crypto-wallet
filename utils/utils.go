package utils

import (
	"github.com/rniveau/crypto-wallet/client"
	"github.com/rniveau/crypto-wallet/model"
)

var clientMongo = client.GetClient()

func GetEuroBudget() *model.Budget {
	euroBudget := clientMongo.GetBudgetByCurrency(model.Euro)
	if euroBudget == nil {
		euro := model.Euro
		euroBudget = &model.Budget{Currency: &euro, Total: float64(0)}
	}
	return euroBudget
}

/*
The percentage is alway positive
 */
func GetPercentageBetweenTwoValue(f1 float64, f2 float64) float64 {
	f3 := f2 / f1;
	if f3 > 1.0 {
		return (f3 - 1) * 100
	}
	return 100 - (f3 * 100)
}

func GetOrderFromOperation(operation *model.Operation) *model.Order {
	if operation.BuyOrder != nil {
		return operation.BuyOrder
	}
	return operation.SellOrder
}