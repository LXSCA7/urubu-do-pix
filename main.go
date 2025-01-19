package main

import (
	"log"
	"urubu-do-pix/config"
	"urubu-do-pix/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// utils.Menu3()
	app := fiber.New()
	routes.SetupRoutes(app)
	config.InitDB()
	log.Fatal(app.Listen(":8000"))
}
