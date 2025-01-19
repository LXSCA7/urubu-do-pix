package models

import (
	"time"
)

type Transaction struct {
	Type     string    `bson:"type"`
	Value    float64   `bson:"value"`
	Date     time.Time `bson:"date"`
	Receiver string    `bson:"receiver,omitempty"`
}
