package sensorlog

import (
	"context"
	"time"

	"github.com/MahdiRazaqi/iotrc/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Sensorlog model
type Sensorlog struct {
	ID           primitive.ObjectID `bson:"_id"`
	Created      time.Time          `bson:"created"`
	TempDHT      float64            `bson:"temp_dht"`
	HumidityDHT  float64            `bson:"humidity_dht"`
	DustHumidity float64            `bson:"dust_humidity"`
	Light        float64            `bson:"light"`
	PompStatus   bool               `bson:"pomp_status"`
	LampStatus   bool               `bson:"lamp_status"`
}

func (l *Sensorlog) collection() *mongo.Collection {
	return database.Mongo.Collection("sensorlog")
}

// InsertOne sensorlog to database
func (l *Sensorlog) InsertOne() error {
	l.ID = primitive.NewObjectID()
	l.Created = time.Now()
	_, err := l.collection().InsertOne(context.Background(), database.ConvertToBson(l))
	return err
}
