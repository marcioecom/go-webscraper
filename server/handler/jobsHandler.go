package handler

import (
	"strconv"
	"webscraper/server/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllJobs(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))

	jobs, err := models.Jobs(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(jobs)
}
