package api

import (
	"net/http"
	"github.com/rniveau/crypto-wallet/client"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rniveau/crypto-wallet/model"
	"strconv"
	"io"
)

func GetBudgetFromCurrency(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	currency, err := strconv.Atoi(params["currency"])
	if err != nil || model.Currency(currency) < model.Bitcoin || model.Currency(currency) > model.End {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Bad currency")
		return
	}
	json.NewEncoder(response).Encode(client.GetBudgetByCurrency(model.Currency(currency)))
}

func GetBudgets(response http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(response).Encode(client.GetBudgets())
}

func AddEuro(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	currency := model.Euro
	money, err := strconv.ParseFloat(params["money"], 64)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Bad money")
		return
	}
	budget := client.GetOrCreateBudget(&currency)
	if money < 0 && budget.Available + money < 0 {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Can't subtract this money from current budget")
		return
	}
	budget.Available += money
	budget.Total += money
	transaction := model.Transaction{Total: money, Date: model.Now()}
	if budget.Transactions == nil {
		budget.Transactions = &[]model.Transaction{}
	}
	transactions := append(*budget.Transactions, transaction)
	budget.Transactions = &transactions
	client.UpsertBudget(budget)
}