package api

import (
	"github.com/rniveau/crypto-wallet/model"
	"gopkg.in/mgo.v2"
	"testing"

	"net/http"
	"net/http/httptest"
	"strings"
)

type mockClientMongo struct {
}

func (mock mockClientMongo) GetSession() *mgo.Session {
	return nil
}
func (mock mockClientMongo) GetOperation(id string) *model.Operation {
	return nil
}
func (mock mockClientMongo) GetCollection(collection string) *mgo.Collection {
	return nil
}
func (mock mockClientMongo) UpsertBudget(budget *model.Budget) {
}
func (mock mockClientMongo) GetChildrenOperation(parentId string) *[]model.Operation {
	return nil

}
func (mock mockClientMongo) GetOperations() []model.Operation {
	return nil

}
func (mock mockClientMongo) GetBudgets() []model.Budget {
	return nil

}
func (mock mockClientMongo) GetBudgetByCurrency(currency model.Currency) *model.Budget {
	return nil

}
func (mock mockClientMongo) GetOrCreateBudget(currency *model.Currency) *model.Budget {
	budget := model.Budget{}
	return &budget

}

var mockMongo = mockClientMongo{}

func TestCreateOperationBitcoinFromEuro(test *testing.T) {
	str := "{\"quantity\": 1, \"currency\": 1, \"description\": \"\", \"buy_order\": {\"price\": 1, \"euro_price\": 1,  \"currency\": 2}}"

	request, _ := http.NewRequest("POST", "test", strings.NewReader(str))
	writer := httptest.NewRecorder()
	clientMongo = &mockMongo
	CreateOperation(writer, request)
}
