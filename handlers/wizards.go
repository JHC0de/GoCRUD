package handlers

import (
	"github.com/JHC0de/GoCRUD/database"
	"github.com/JHC0de/GoCRUD/models"
	"github.com/gofiber/fiber/v2"
)

func ListWizards(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateWizard(c *fiber.Ctx) error {
	fact := new(models.Wizard)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func RemoveWizard(c *fiber.Ctx) error {
	id := c.Params("id")
	var fact models.Wizard

	result := database.DB.Db.Delete(&fact, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
