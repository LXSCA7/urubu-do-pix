package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"urubu-do-pix/config"
	"urubu-do-pix/middleware"
	"urubu-do-pix/models"
	"urubu-do-pix/services"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func Deposit(c fiber.Ctx) error {
	var depositRequest struct {
		Username string  `json:"username"`
		Amount   float64 `json:"amount"`
	}

	err := json.Unmarshal(c.Body(), &depositRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"expected": depositRequest,
		})
	}

	if depositRequest.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "quantidade deve ser maior que 0",
		})
	}

	user := services.GetByUsername(depositRequest.Username)
	if user.Id.IsZero() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"response": "user not found",
		})
	}

	err = services.AddUserBalance(&user, depositRequest.Amount, "Deposito")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// user returns just for development
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": "quantia adicionada",
		"user":     user,
	})
}

func Withdraw(c fiber.Ctx) error {
	middleware.Verify(c)
	username := c.Locals("username")
	if username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username not found in request context.",
		})
	}
	usernameStr, ok := username.(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username is not valid in request context.",
		})
	}

	var body struct {
		Amount float64 `json:"amount"`
	}

	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if body.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid amount",
		})
	}

	return services.Withdraw(c, usernameStr, body.Amount)
}

func Transfer(c fiber.Ctx) error {
	senderUsername, err := Authenticate(c)
	if err != nil {
		return err
	}

	var body struct {
		Username string  `json:"username"`
		Amount   float64 `json:"amount"`
	}

	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":   "error",
			"error":    err.Error(),
			"expected": body,
		})
	}

	if body.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "amount cannot be less than 0",
		})
	}

	recipient := services.GetByUsername(body.Username)
	if recipient.Id.IsZero() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"response": "user not found",
		})
	}

	return services.Transfer(c, senderUsername, &recipient, body.Amount)
}

func CreateDepositsFieldForAllUsers() error {
	collection := config.GetCollection("urubu_users")

	filter := bson.M{
		"deposits": bson.M{"$exists": false},
	}
	update := bson.M{
		"$set": bson.M{"deposits": []models.Deposit{}}, // Array vazio de depósitos
	}

	result, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to create deposits field: %v", err)
	}

	fmt.Printf("Campo 'deposits' criado para %d usuarios\n", result.ModifiedCount)
	return nil
}
