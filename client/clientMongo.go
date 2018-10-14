package client 

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/rniveau/crypto-wallet/model"
)

const (
	DB string = "test"
	OperationCollection string = "operations"
	BudgetCollection string = "budgets"
)

type ClientMongo interface { 
	GetSession() *mgo.Session
	GetOperation() *interface{}
	GetCollection(collection string) *mgo.Collection
	UpsertBudget(budget *model.Budget)
	GetChildrenOperation(parentId string) *[]model.Operation
	GetOperations() []model.Operation
	GetBudgets() []model.Budget
	GetBudgetByCurrency(currency model.Currency) *model.Budget
	GetOrCreateBudget(currency *model.Currency) *model.Budget
}

type clientMongo struct {

	session *mgo.Session

}

var (
	client = &clientMongo{}
)

func GetClient() *clientMongo {
	return client
}

func (c clientMongo) getSession() *mgo.Session  {
	if c.session == nil {
		log.Println("Init session mongodb")
		c.session, _ = mgo.Dial("localhost")
	}
	return c.session
}

func GetSession() *mgo.Session {
	return client.getSession()
}

func (client clientMongo) getCollection(collection string) *mgo.Collection {
	return client.getSession().DB(DB).C(collection)
}

func (client clientMongo) _getById(id string, collection string) *interface{} {
	if bson.IsObjectIdHex(id) {
		var value *interface{}
		client.getCollection(collection).FindId(bson.ObjectIdHex(id)).One(&value)
		return value
	} else {
		return nil
	}
}

func (client clientMongo) getOperation(id string) *model.Operation {
	var operation *model.Operation
	interf := client._getById(id, OperationCollection)
	if interf == nil {
		return nil
	}
	bsonM := (*(interf)).(bson.M)
	bsonBytes, err := bson.Marshal(bsonM)
	if err != nil {
		log.Println(err)
		return nil
	}
	bson.Unmarshal(bsonBytes, &operation)
	return operation
}

func (client clientMongo) getBudget(id string) *interface{} {
	return client._getById(id, BudgetCollection)
}

func (client clientMongo) GetCollection(collection string) *mgo.Collection {
	return client.getCollection(collection)
}

func (client clientMongo) GetOperation(id string) *model.Operation {
	return client.getOperation(id)
}

func (client clientMongo) GetChildrenOperation(parentId string) *[]model.Operation {
	var values []model.Operation
	err := client.getCollection(OperationCollection).Find(bson.M{"parentid": parentId}).All(&values)
	if err != nil {
		log.Println(err)
	}
	return &values
}

func (client clientMongo) GetOperations() []model.Operation {
	var values []model.Operation
	err := client.getCollection(OperationCollection).Find(bson.M{}).All(&values)
	if err != nil {
		log.Println(err)
	}
	return values
}

func (client clientMongo) GetBudgets() []model.Budget {
	var values []model.Budget
	client.getCollection(BudgetCollection).Find(bson.M{}).All(&values)
	return values
}

func (client clientMongo) GetBudgetByCurrency(currency model.Currency) *model.Budget {
	var budget *model.Budget
	client.getCollection(BudgetCollection).Find(bson.M{"currency": currency}).One(&budget)
	return budget
}

func (client clientMongo) GetOrCreateBudget(currency *model.Currency) *model.Budget {
	currencyBudget := client.GetBudgetByCurrency(*currency)
	if currencyBudget == nil {
		currencyBudget = &model.Budget{Currency: currency, Total: float64(0)}
	}
	return currencyBudget
}

func (client clientMongo) UpsertBudget(budget *model.Budget) {
	if !bson.IsObjectIdHex(budget.Id.Hex()) {
		budget.Id = bson.NewObjectId()
	}
	selector := bson.M{"_id": budget.Id}
	if _, err := client.getCollection(BudgetCollection).Upsert(selector, budget); err != nil {
		log.Println(err)
	}
}