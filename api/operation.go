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
	budget := client.GetOrCreateBudget(operation.Currency)
	var euroBudget *model.Budget
	if *order.Currency != model.Euro {
		euroBudget = utils.GetEuroBudget()
		if operation.BuyOrder != nil {
			if euroBudget.Available - order.EuroPrice < 0 {
				response.WriteHeader(http.StatusBadRequest)
				io.WriteString(response, "No enough euro budget")
				return
			}
			euroBudget.Available -= order.EuroPrice
		} else {
			euroBudget.Available += order.EuroPrice
			if euroBudget.Available > euroBudget.Total {
				if euroBudget.Transactions == nil {
					euroBudget.Transactions = &[]model.Transaction{}
				}
				transactions := append(*euroBudget.Transactions, model.Transaction{Date: model.Now(), Total: euroBudget.Available - euroBudget.Total})
				euroBudget.Transactions = &transactions
				euroBudget.Total = euroBudget.Available
			}
		}
	}
	if operation.BuyOrder != nil {
		budget.Total += operation.Quantity
		budget.Available += operation.Quantity
		currencyBudget := client.GetOrCreateBudget(operation.BuyOrder.Currency)
		if currencyBudget.Available - order.Price < 0 {
			response.WriteHeader(http.StatusBadRequest)
			io.WriteString(response, "No enough currency budget")
			return
		}
		currencyBudget.Available -= order.Price
		client.UpsertBudget(currencyBudget)
	} else {
		budget.Total -= operation.Quantity
		budget.Available -= operation.Quantity
		currencyBudget := client.GetOrCreateBudget(operation.SellOrder.Currency)
		currencyBudget.Available += order.Price
		if currencyBudget.Available > currencyBudget.Total {
			if *currencyBudget.Currency == model.Euro {
				if currencyBudget.Transactions == nil {
					currencyBudget.Transactions = &[]model.Transaction{}
				}
				transactions := append(*currencyBudget.Transactions, model.Transaction{Date: model.Now(), Total: currencyBudget.Available - currencyBudget.Total})
				currencyBudget.Transactions = &transactions
			}
			currencyBudget.Total = currencyBudget.Available
		}
		client.UpsertBudget(currencyBudget)
	}
	if euroBudget != nil {
		client.UpsertBudget(euroBudget)
	}
	client.UpsertBudget(budget)
	client.GetCollection(client.OperationCollection).Insert(operation)
}
