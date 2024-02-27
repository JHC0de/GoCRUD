// cmd/routes.go

package main

import (
	"github.com/JHC0de/GoCRUD/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupFactRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/list/facts", handlers.ListFacts)
	app.Post("/create/fact", handlers.CreateFact)
	app.Delete("/remove/fact/:id", handlers.RemoveFact)
}

func setupWizardRoutes(app *fiber.App) {
	app.Get("/list/wizards", handlers.ListWizards)
	app.Post("/create/wizard", handlers.CreateWizard)
	app.Delete("/remove/wizard/:id", handlers.RemoveWizard)
	app.Patch("/damage/wizard/:id/:damage", handlers.DamageWizard)
}
