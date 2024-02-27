package handlers

import (
	"github.com/JHC0de/GoCRUD/database"
	"github.com/JHC0de/GoCRUD/models"
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func ListWizards(c *fiber.Ctx) error {
	wizards := []models.Wizard{}
	database.DB.Db.Find(&wizards)

	return c.Status(200).JSON(wizards)
}

func CreateWizard(c *fiber.Ctx) error {
	wizard := new(models.Wizard)
	if err := c.BodyParser(wizard); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&wizard)

	return c.Status(200).JSON(wizard)
}

func RemoveWizard(c *fiber.Ctx) error {
	id := c.Params("id")
	var wizard models.Wizard

	result := database.DB.Db.Delete(&wizard, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func DamageWizard(c *fiber.Ctx) error {
	id := c.Params("id")
	damage := c.Params("damage")

	var wizard models.Wizard
	database.DB.Db.First(&wizard, id)
	result := database.DB.Db.Model(&wizard).Update(models.Wizard{
		Health: 12
	})
	log.Print(result)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
