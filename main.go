package main

import (
	"io"
	"log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"github.com/rniveau/crypto-wallet/client"
	"github.com/rniveau/crypto-wallet/model"
)

func httpHandler(response http.ResponseWriter, request *http.Request) {
    response.WriteHeader(http.StatusNoContent)
}

func getOperation(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    operation := client.GetOperation(params["id"])
	json.NewEncoder(response).Encode(*operation)
}

func createOperation(response http.ResponseWriter, request *http.Request) {
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

func main() {
//	c := client.GetCollection("test")
//	c.Insert(model.Operation{Currency:model.Euro})
  	router := mux.NewRouter()
    router.HandleFunc("/", httpHandler)
    router.HandleFunc("/api", httpHandler).Methods("GET")
    router.HandleFunc("/api/operation/{id}", getOperation).Methods("GET")
    router.HandleFunc("/api/operation", createOperation).Methods("POST")
    router.HandleFunc("/api/operations", getOperation).Methods("GET")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Error to start http", err)
	}
}