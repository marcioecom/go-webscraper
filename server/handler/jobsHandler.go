package handler

import (
	"webscraper/server/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllJobs(c *fiber.Ctx) error {
	jobs, err := models.Jobs()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(jobs)
}
