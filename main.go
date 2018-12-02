package main

import (
	"github.com/rniveau/crypto-wallet/config"
	"github.com/spf13/viper"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rniveau/crypto-wallet/api"
)

func httpHandler(response http.ResponseWriter, request *http.Request) {
    response.WriteHeader(http.StatusNoContent)
}

func configuration() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/crypto-wallet/")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatalln("Fatal error config file: ", err)
	}
	err = viper.Unmarshal(&config.GlobalConfig)
	if err != nil { // Handle errors reading the config file
		log.Fatalln("Fatal error config file: ", err)
	}
}

func main() {
	configuration()
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