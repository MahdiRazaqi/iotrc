package database

import (
	"context"

	"github.com/MahdiRazaqi/iotrc/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo connection
var Mongo *mongo.Database

// Connect to MongoDB
func Connect() error {
	cfg := config.Config.Mongo

	opt := options.Client()
	if cfg.Host != "" {
		opt.ApplyURI(cfg.Host)
	}

	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		return err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return err
	}
	Mongo = client.Database(cfg.DB)

	return nil
}

// ConvertToBson convert interface to bson object
func ConvertToBson(d interface{}) bson.M {
	val, _ := bson.Marshal(d)
	data := new(bson.M)
	bson.Unmarshal(val, data)
	return *data
}
