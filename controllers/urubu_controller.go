package controllers

import (
	"encoding/json"
	"urubu-do-pix/middleware"
	"urubu-do-pix/services"

	"github.com/gofiber/fiber/v3"
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
