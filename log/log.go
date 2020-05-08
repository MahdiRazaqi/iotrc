package log

import "go.mongodb.org/mongo-driver/bson/primitive"

// Log model
type Log struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id"`
	Light float64            `bson:"light" json:"light"`
}
