package handler

import (
	"webscraper/server/models"
	"webscraper/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly models.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}

	err = models.CreateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error could not create goly in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(goly)
}
