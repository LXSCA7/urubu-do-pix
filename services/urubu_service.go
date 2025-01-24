package services

import (
	"context"
	"fmt"
	"time"
	"urubu-do-pix/config"
	"urubu-do-pix/models"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUserBalance(user *models.User, amount float64, transactionType string) error {
	user.Balance += amount
	newTransaction := models.Transaction{
		Value: amount,
		Type:  transactionType,
		Date:  time.Now(),
	}
	user.Transactions = append(user.Transactions, newTransaction)

	return UpdateUserBalance(user, newTransaction)
}

func RemoveUserBalance(user *models.User, amount float64, transactionType string) error {
	user.Balance -= amount
	newTransaction := models.Transaction{
		Value: amount,
		Type:  transactionType,
		Date:  time.Now(),
	}
	user.Transactions = append(user.Transactions, newTransaction)

	return UpdateUserBalance(user, newTransaction)
}

func UpdateUserBalance(user *models.User, newTransaction models.Transaction) error {
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
		return fmt.Errorf("failed to record a transaction: %v", err)
	}

	return nil
}

func Withdraw(c fiber.Ctx, username string, amount float64) error {
	user := GetByUsername(username)
	if user.Id.IsZero() {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if user.Balance < amount {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid amount",
		})
	}

	RemoveUserBalance(&user, amount, "Saque")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func Transfer(c fiber.Ctx, senderUsername string, receiver *models.User, amount float64) error {
	sender := GetByUsername(senderUsername)
	if sender.Id.IsZero() {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if sender.Balance < amount {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "insufficient funds",
		})
	}

	senderMsg := "transferencia para" + receiver.Username
	receiverMsg := "transferencia de" + sender.Username

	err := RemoveUserBalance(&sender, amount, senderMsg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = AddUserBalance(receiver, amount, receiverMsg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// err := UpdateUserBalance(&sender, senderMsg)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"receiver": receiver,
		"sender":   sender,
	})
}
