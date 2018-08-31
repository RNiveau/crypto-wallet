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
}

type clientMongo struct {

	session *mgo.Session

}

var (
	client *clientMongo = &clientMongo{}
)

func (c clientMongo) getClient() *mgo.Session  {
	if c.session == nil {
		log.Println("Init session mongodb")
		c.session, _ = mgo.Dial("localhost")
	}
	return c.session
}

func GetSession() *mgo.Session {
	return client.getClient()
}

func (client clientMongo) getCollection(collection string) *mgo.Collection {
	return client.getClient().DB(DB).C(collection)
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
	bsonM := (*(client._getById(id, OperationCollection))).(bson.M)
	bsonBytes, _ := bson.Marshal(bsonM)
	bson.Unmarshal(bsonBytes, &operation)
	return operation
}

func (client clientMongo) getBudget(id string) *interface{} {
	return client._getById(id, BudgetCollection)
}

func GetCollection(collection string) *mgo.Collection {
	return client.getCollection(collection)
}

func GetOperation(id string) *model.Operation {
	return client.getOperation(id)
}

func GetOperations() []model.Operation {
	var values []model.Operation
	err := client.getCollection(OperationCollection).Find(bson.M{}).All(&values)
	if err != nil {
		log.Println(err)
	}
	return values
}

func GetBudget(id string) *interface{} {
	return client.getBudget(id)
}

func GetBudgetByCurrency(currency model.Currency) *model.Budget {
	var budget *model.Budget
	client.getCollection(BudgetCollection).Find(bson.M{"currency": currency}).One(&budget)
	return budget
}

func UpsertBudget(budget *model.Budget) {
	client.getCollection(BudgetCollection).Upsert(bson.M{}, budget)
}