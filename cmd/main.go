package main

import (
	"github.com/JHC0de/GoCRUD/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupFactRoutes(app)
	setupWizardRoutes(app)

	app.Listen(":3000")
}
