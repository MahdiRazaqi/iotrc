package log

import (
	"context"
	"time"

	"github.com/MahdiRazaqi/iotrc/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Log model
type Log struct {
	ID           primitive.ObjectID `bson:"_id"`
	Created      time.Time          `bson:"created"`
	TempDHT      int                `bson:"temp_dht"`
	HumidityDHT  int                `bson:"humidity_dht"`
	DustHumidity int                `bson:"dust_humidity"`
	Light        int                `bson:"light"`
	PompStatus   bool               `bson:"pomp_status"`
	LampStatus   bool               `bson:"lamp_status"`
}

func (l *Log) collection() *mongo.Collection {
	return database.Mongo.Collection("log")
}

// InsertOne log to database
func (l *Log) InsertOne() error {
	l.ID = primitive.NewObjectID()
	l.Created = time.Now()
	_, err := l.collection().InsertOne(context.Background(), database.ConvertToBson(l))
	return err
}
