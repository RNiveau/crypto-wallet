package api

import (
	"io"
	"log"
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"

	"github.com/rniveau/crypto-wallet/model"
	"github.com/rniveau/crypto-wallet/client"
	"github.com/rniveau/crypto-wallet/utils"
)

func GetOperation(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	operation := client.GetOperation(params["id"])
	json.NewEncoder(response).Encode(*operation)
}

func GetOperations(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	operations := client.GetOperations()
	json.NewEncoder(response).Encode(operations)
}

func CreateOperation(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var operation model.Operation
	err := decoder.Decode(&operation)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Can't decode json")
		return
	}
	operation.Id = bson.NewObjectId()
	if operation.ParentId != "" {
		operation.Parent = client.GetOperation(operation.ParentId)
	}
	if err = operation.Valid(); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, err.Error())
		return
	}
	order := utils.GetOrderFromOperation(&operation)
	budget := client.GetBudgetByCurrency(*operation.Currency)
	if budget == nil {
		budget = &model.Budget{Currency: operation.Currency, Total: float64(0)}
	}
	if *order.Currency != model.Euro {
		euroBudget := utils.GetEuroBudget()
		if operation.BuyOrder != nil {
			if euroBudget.Available - order.EuroPrice < 0 {
				response.WriteHeader(http.StatusBadRequest)
				io.WriteString(response, "No enough euro budget")
				return
			}
			euroBudget.Available -= order.EuroPrice
		} else {
			euroBudget.Available += order.EuroPrice
		}
		client.UpsertBudget(euroBudget)
	}
	if operation.BuyOrder != nil {
		budget.Total += operation.Quantity
		budget.Available += operation.Quantity
		currencyBudget := client.GetBudgetByCurrency(*operation.BuyOrder.Currency)
		if currencyBudget == nil {
			currencyBudget = &model.Budget{Currency: operation.BuyOrder.Currency, Total: float64(0)}
		}
		if currencyBudget.Available - order.Price < 0 {
			response.WriteHeader(http.StatusBadRequest)
			io.WriteString(response, "No enough currency budget")
			return
		}
		currencyBudget.Available -= order.Price
		client.UpsertBudget(currencyBudget)
	} else {

	}
	client.UpsertBudget(budget)
	client.GetCollection(client.OperationCollection).Insert(operation)
}
