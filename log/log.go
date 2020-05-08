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
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	Created time.Time          `bson:"created" json:"created"`
	Light   float64            `bson:"light" json:"light"`
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
