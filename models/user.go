package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Password     string             `bson:"password"`
	Balance      float64            `bson:"balance"`
	Transactions []Transaction      `bson:"transactions"`
	Deposits     []Deposit          `bson:"deposits"`
}
