// cmd/routes.go

package main

import (
	"github.com/JHC0de/GoCRUD/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
