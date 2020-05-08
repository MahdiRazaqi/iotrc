package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database

// Connect to MongoDB
func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("connecting to mongo ", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("connecting to mongo ", err)
	}

	Mongo = client.Database("iotrc")
}

// ConvertToBson convert interface to bson object
func ConvertToBson(d interface{}) bson.M {
	val, _ := bson.Marshal(d)
	data := new(bson.M)
	bson.Unmarshal(val, data)
	return *data
}
