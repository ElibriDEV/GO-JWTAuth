package initializators

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoManager MongoManagerStruct

type MongoManagerStruct struct {
	Collection *mongo.Collection
}

func (manager *MongoManagerStruct) Connect(mongoUri string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal("Mongo connection not established:\n", err.Error())
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Mongo connection not established:\n", err.Error())
	}
	collection := client.Database("auth").Collection("token")
	manager.Collection = collection
}

func MongoInit(mongoUri string) {
	MongoManager.Connect(mongoUri)
	return
}
