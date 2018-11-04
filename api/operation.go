package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"github.com/rniveau/crypto-wallet/client"
	"github.com/rniveau/crypto-wallet/model"
	"github.com/rniveau/crypto-wallet/utils"
)

var clientMongo = client.GetClient()

func GetOperation(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	operation := clientMongo.GetOperation(params["id"])
	operation.Children = clientMongo.GetChildrenOperation(params["id"])
	json.NewEncoder(response).Encode(*operation)
}

func GetOperations(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	operations := clientMongo.GetOperations()
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
		operation.Parent = clientMongo.GetOperation(operation.ParentId)
	}
	if err = operation.Valid(); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, err.Error())
		return
	}
	order := utils.GetOrderFromOperation(&operation)
	budget := clientMongo.GetOrCreateBudget(operation.Currency)
	if operation.BuyOrder != nil {
		currencyBudget := clientMongo.GetOrCreateBudget(operation.BuyOrder.Currency)
		if currencyBudget.Available - order.Price < 0 {
			response.WriteHeader(http.StatusBadRequest)
			io.WriteString(response, "No enough currency budget")
			return
		}
		budget.Total += operation.Quantity
		budget.Available += operation.Quantity
		currencyBudget.Available -= order.Price
		clientMongo.UpsertBudget(currencyBudget)
	} else {
		budget.Total -= operation.Quantity
		budget.Available -= operation.Quantity
		currencyBudget := clientMongo.GetOrCreateBudget(operation.SellOrder.Currency)
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
		clientMongo.UpsertBudget(currencyBudget)
	}
	clientMongo.UpsertBudget(budget)
	clientMongo.InsertOperation(&operation)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(operation)
}
