package api

import (
	"encoding/json"
	"github.com/rniveau/crypto-wallet/model"
	"net/http"
)

func GetCryptos(response http.ResponseWriter, request *http.Request) {
	maps := make(map[string]model.Currency)
	maps["euro"] = model.Euro
	maps["bitcoin"] = model.Bitcoin
	maps["ether"] = model.Ether
	maps["ripple"] = model.Ripple
	maps["iost"] = model.IOST
	json.NewEncoder(response).Encode(maps)


}
