package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rniveau/crypto-wallet/api"
)

func httpHandler(response http.ResponseWriter, request *http.Request) {
    response.WriteHeader(http.StatusNoContent)
}


func main() {
//	c := client.GetCollection("test")
//	c.Insert(model.Operation{Currency:model.Euro})
  	router := mux.NewRouter()
    router.HandleFunc("/", httpHandler)
    router.HandleFunc("/api", httpHandler).Methods("GET")
    router.HandleFunc("/api/operation/{id}", api.GetOperation).Methods("GET")
    router.HandleFunc("/api/operation", api.CreateOperation).Methods("POST")
    router.HandleFunc("/api/operations", api.GetOperations).Methods("GET")
    router.HandleFunc("/api/cryptos", api.GetCryptos).Methods("GET")
    router.HandleFunc("/api/budgets", api.GetBudgets).Methods("GET")
    router.HandleFunc("/api/budget/{currency}", api.GetBudgetFromCurrency).Methods("GET")
    router.HandleFunc("/api/budget/euro/{money}", api.AddEuro).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Error to start http", err)
	}
}