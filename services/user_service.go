package services

import (
	"context"
	"fmt"
	"os"
	"time"
	"urubu-do-pix/config"
	"urubu-do-pix/models"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateItem(user *models.User, collectionName string) error {
	collection := config.GetCollection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	primitiveId, _ := result.InsertedID.(primitive.ObjectID)
	user.Id = primitiveId
	return nil
}

func GetByUsername(username string) models.User {
	collection := config.GetCollection("urubu_users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result *models.User
	err := collection.FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&result)
	if err != nil {
		return models.User{}
	}
	return *result
}

func Login(c fiber.Ctx, user models.User) error {
	dbUser := GetByUsername(user.Username)

	if dbUser.Id.IsZero() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"response": "user not found",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"response": "wrong password",
		})
	}
	exp := time.Now().Add(24 * time.Hour).Unix()
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "logged id",
		"token":   t,
		"expires": time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
}

func UpdateEmpty(c fiber.Ctx) error {
	// Obtém a coleção do MongoDB
	collection := config.GetCollection("urubu_users")

	// Filtro para encontrar documentos onde o campo "transactions" é null
	filter := bson.M{"transactions": nil}

	// Atualização para definir "transactions" como um array vazio
	update := bson.M{
		"$set": bson.M{"transactions": []models.Transaction{}}, // Define como um array vazio
	}

	// Executa a operação de atualização
	_, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("falha ao atualizar documentos: %w", err)
	}

	return nil
}
