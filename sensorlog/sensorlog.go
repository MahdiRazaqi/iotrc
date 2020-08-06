package sensorlog

import (
	"context"
	"time"

	"github.com/MahdiRazaqi/iotrc/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Sensorlog model
type Sensorlog struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Created      time.Time          `bson:"created" json:"created"`
	TempDHT      float64            `bson:"temp_dht" json:"temp_dht"`
	HumidityDHT  float64            `bson:"humidity_dht" json:"humidity_dht"`
	DustHumidity float64            `bson:"dust_humidity" json:"dust_humidity"`
	Light        float64            `bson:"light" json:"light"`
	MQ9          float64            `bson:"mq9" json:"mq9"`
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

// Find sensorlog
func Find(filter bson.M, page, limit int) ([]Sensorlog, error) {
	s := new(Sensorlog)
	ctx := context.Background()
	options := options.Find()
	options.SetLimit(int64(limit))
	if page > 0 {
		options.SetSkip(int64((page - 1) * limit))
	}

	cursor, err := s.collection().Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	sensorlogs := []Sensorlog{}
	for cursor.Next(ctx) {
		s := new(Sensorlog)
		if err := cursor.Decode(s); err != nil {
			continue
		}
		sensorlogs = append(sensorlogs, *s)
	}

	return sensorlogs, nil
}
