package auth

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jwt-auth/initializators"
	"log"
)

type UserToken struct {
	ID      primitive.ObjectID `bson:"_id"`
	Guid    string
	Refresh string
}

type UserTokenRepository struct{}

func (repository *UserTokenRepository) Insert(guid, refreshToken string, ch chan interface{}) {
	result, err := initializators.MongoManager.Collection.UpdateOne(
		context.TODO(),
		bson.D{{"guid", guid}},
		bson.D{{"$set", bson.D{{"refresh", Encode(refreshToken)}}}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	ch <- result.UpsertedID
}

func (repository *UserTokenRepository) GetOne(guid string, ch chan *UserToken) {
	var result UserToken
	err := initializators.MongoManager.Collection.FindOne(
		context.TODO(),
		bson.D{{"guid", guid}},
	).Decode(&result)
	if err != nil {
		log.Fatal(err.Error())
	}
	println(guid)
	println(result.Guid)
	ch <- &result
}
