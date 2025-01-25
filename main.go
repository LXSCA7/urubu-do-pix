package main

import (
	"log"
	"urubu-do-pix/config"
	"urubu-do-pix/routes"
	"urubu-do-pix/routine"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// utils.Menu3()
	config.InitDB()
	go routine.Start()

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
