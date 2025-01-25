package controllers

import (
	"encoding/json"
	"urubu-do-pix/middleware"
	"urubu-do-pix/models"
	"urubu-do-pix/services"
	"urubu-do-pix/utils"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c fiber.Ctx) error {
	var newUser models.User
	err := json.Unmarshal(c.Body(), &newUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"result":   err.Error(),
			"expected": models.User{},
		})
	}
	errorList := utils.IsPasswordStrong(newUser.Password)
	if len(errorList) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errorList,
		})
	}

	existingUser := services.GetByUsername(newUser.Username)
	if !existingUser.Id.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"response": "username already exists",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	newUser.Password = string(hash)
	newUser.Balance = 0.0
	newUser.Transactions = []models.Transaction{}

	services.CreateItem(&newUser, "urubu_users")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"response":      "usuario cadastrado com sucesso.",
		"user_id":       newUser.Id,
		"user_username": newUser.Username,
	})
}

func GetInfo(c fiber.Ctx) error {
	username := c.Params("username")
	user := services.GetByUsername(username)
	if user.Id.IsZero() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"response": "user not found",
		})
	}

	return c.JSON(user)
}

func Login(c fiber.Ctx) error {
	var user models.User
	json.Unmarshal(c.Body(), &user)
	return services.Login(c, user)
	// return c.Status(fiber.StatusOK).JSON(dbUser)
}

func Authenticate(c fiber.Ctx) (string, error) {
	middleware.Verify(c)
	username := c.Locals("username")
	if username == nil {
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username not found in request context.",
		})
	}

	usernameStr, ok := username.(string)
	if !ok {
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username is not valid in request context.",
		})
	}

	return usernameStr, nil
}

func UpdateEmpty(c fiber.Ctx) error {
	return services.UpdateEmpty(c)
}
