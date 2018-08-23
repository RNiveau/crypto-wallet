package main

import (
	"log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rniveau/crypto-wallet/model"
)

func httpHandler(w http.ResponseWriter, request *http.Request) {
    w.WriteHeader(http.StatusNoContent)
}

func getOperation(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    //params := mux.Vars(request)
    operation := model.Operation{Id:"test", Currency: model.Euro}
	json.NewEncoder(w).Encode(operation)
}

func main() {
  	router := mux.NewRouter()
    router.HandleFunc("/", httpHandler)
    router.HandleFunc("/api", httpHandler).Methods("GET")
    router.HandleFunc("/api/operation/{id}", getOperation).Methods("GET")
    router.HandleFunc("/api/operations", getOperation).Methods("GET")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Error to start http", err)
	}
}