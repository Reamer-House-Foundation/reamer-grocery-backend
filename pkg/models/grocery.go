package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Grocery struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Quantity int32              `bson:"quantity"`
}
