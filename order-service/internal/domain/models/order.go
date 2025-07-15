package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Order struct {
	Id     bson.ObjectID `bson:"_id,omitempty"`
	BookId int64         `bson:"book_id"`
	Status string        `bson:"status"`
	Price  float64       `bson:"price"`
	Place  string        `bson:"place"`
}
