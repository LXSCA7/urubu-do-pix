package routes

import (
	"urubu-do-pix/controllers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	// app.Group("/api")
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("hello, world 👋")
	})

	app.Post("/user/create", func(c fiber.Ctx) error {
		return controllers.CreateUser(c)
	})

	app.Post("user/login", func(c fiber.Ctx) error {
		return controllers.Login(c)
	})

	app.Get("user/:username", func(c fiber.Ctx) error {
		return controllers.GetInfo(c)
	})

	// urubu:
	app.Post("urubu/deposit", func(c fiber.Ctx) error {
		return controllers.Deposit(c)
	})

	app.Post("urubu/withdraw", func(c fiber.Ctx) error {
		return controllers.Withdraw(c)
	})

	app.Post("urubu/transfer", func(c fiber.Ctx) error {
		return controllers.Transfer(c)
	})
}
