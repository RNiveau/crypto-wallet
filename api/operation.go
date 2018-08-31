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
	if err = operation.Valid(); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, err.Error())
		return
	}
	client.GetCollection(client.OperationCollection).Insert(operation)
	log.Println(operation.Id)
}
