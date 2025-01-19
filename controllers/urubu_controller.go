package controllers

import (
	"encoding/json"
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

	err = services.UpdateBalance(&user, depositRequest.Amount)
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
