package handlers

import "github.com/gofiber/fiber/v2"

func GetAllJobs(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "All jobs",
	})
}
