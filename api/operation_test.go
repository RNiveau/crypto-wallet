package api

import (
	"testing"

	"net/http"
	"strings"
	"net/http/httptest"
)

func TestCreateOperationBitcoinFromEuro(test *testing.T) {
	str := "{\"quantity\": 1, \"currency\": 1, \"description\": \"\", \"buy_order\": {\"price\": 1, \"euro_price\": 1,  \"currency\": 2}}"

	request, _ := http.NewRequest("POST", "test", strings.NewReader(str))
	writer := httptest.NewRecorder()
	CreateOperation(writer, request)
}

