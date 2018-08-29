package client 

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB string = "test"
	OperationCollection string = "operations"
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

func (client clientMongo) getOperation(id string) *interface{} {
	if bson.IsObjectIdHex(id) {
		var operation *interface{}
		client.getCollection(OperationCollection).FindId(bson.ObjectIdHex(id)).One(&operation)
		return operation
	} else {
		return nil		
	}
}

func GetCollection(collection string) *mgo.Collection {
	return client.getCollection(collection)
}

func GetOperation(id string) *interface{} {
	return client.getOperation(id)
}