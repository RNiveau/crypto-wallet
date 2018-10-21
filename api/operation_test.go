package api

import (
	"github.com/rniveau/crypto-wallet/model"
	"gopkg.in/mgo.v2"
	"testing"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"strings"
)

type mockClientMongo struct {
	budget *model.Budget
	euroBudget *model.Budget
}

func (mock mockClientMongo) GetSession() *mgo.Session {
	return nil
}
func (mock mockClientMongo) GetOperation(id string) *model.Operation {
	return nil
}
func (mock mockClientMongo) GetCollection(collection string) *mgo.Collection {
	collec := mgo.Collection{}
	return &collec
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
	if *currency == model.Euro {
		return mock.euroBudget
	}
	return mock.budget
}

func (mock mockClientMongo) InsertOperation(operation *model.Operation) {
}

var mockMongo = mockClientMongo{}

func TestCreateBuyOperationBitcoinFromEuro(test *testing.T) {
	str := "{\"quantity\": 1, \"currency\": 1, \"description\": \"\", \"buy_order\": {\"price\": 1, \"euro_price\": 1,  \"currency\": 2}}"

	request, _ := http.NewRequest("POST", "test", strings.NewReader(str))
	writer := httptest.NewRecorder()
	clientMongo = &mockMongo
	currency := model.Euro
	euroBudget := model.Budget{Currency: &currency, Total: 10, Available: 10}
	mockMongo.euroBudget = &euroBudget
	bitcoin := model.Bitcoin
	bitcoinBudget := model.Budget{Currency: &bitcoin, Total: 0, Available: 0}
	mockMongo.budget = &bitcoinBudget
	CreateOperation(writer, request)
	assert.Equal(test, float64(9), euroBudget.Available)
	assert.Equal(test, float64(1), bitcoinBudget.Available)
	assert.Equal(test, float64(1), bitcoinBudget.Total)
}
