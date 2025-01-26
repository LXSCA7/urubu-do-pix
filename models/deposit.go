package models

import (
	"time"
)

type Deposit struct {
	Value float64   `bson:"value"`
	Date  time.Time `bson:"date"`
}
