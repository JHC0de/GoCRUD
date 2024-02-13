// cmd/routes.go

package main

import (
	"github.com/JHC0de/GoCRUD/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/list", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

	app.Delete("/remove/:id", handlers.RemoveFact)
}
