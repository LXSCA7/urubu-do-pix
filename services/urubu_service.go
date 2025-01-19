package services

import (
	"context"
	"fmt"
	"time"
	"urubu-do-pix/config"
	"urubu-do-pix/models"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateBalance(user *models.User, amount float64) error {
	user.Balance += amount
	newTransaction := models.Transaction{
		Value: amount,
		Type:  "Deposito",
		Date:  time.Now(),
	}
	user.Transactions = append(user.Transactions, newTransaction)

	collection := config.GetCollection("urubu_users")
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "balance", Value: user.Balance},
		}},
		{Key: "$push", Value: bson.D{
			{Key: "transactions", Value: newTransaction},
		}},
	}

	_, err := collection.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: user.Id}}, update)
	if err != nil {
		return fmt.Errorf("failed to record a deposit: %v", err)
	}

	return nil
}
