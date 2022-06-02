package routes

import (
	"webscraper/server/crawler"
	"webscraper/server/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupJobsRoutes(app *fiber.Router) {
	api := (*app).Group("/jobs")

	api.Get("/", handler.GetAllJobs)
	api.Get("/a", func(c *fiber.Ctx) error {
		jobs := crawler.ScrapJobs()

		return c.Status(fiber.StatusOK).JSON(jobs)
	})
}
