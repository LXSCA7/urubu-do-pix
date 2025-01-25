package routes

import (
	"urubu-do-pix/controllers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("hello, world 👋")
	})

	api.Get("/update-users", func(c fiber.Ctx) error {
		return controllers.UpdateEmpty(c)
	})

	api.Post("/user/create", func(c fiber.Ctx) error {
		return controllers.CreateUser(c)
	})

	api.Post("user/login", func(c fiber.Ctx) error {
		return controllers.Login(c)
	})

	api.Get("user/:username", func(c fiber.Ctx) error {
		return controllers.GetInfo(c)
	})

	// urubu:
	api.Post("urubu/deposit", func(c fiber.Ctx) error {
		return controllers.Deposit(c)
	})

	api.Post("urubu/withdraw", func(c fiber.Ctx) error {
		return controllers.Withdraw(c)
	})

	api.Post("urubu/transfer", func(c fiber.Ctx) error {
		return controllers.Transfer(c)
	})
}
