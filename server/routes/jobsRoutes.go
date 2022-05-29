package routes

import (
	"webscraper/server/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupJobsRoutes(app *fiber.Router) {
	api := (*app).Group("/jobs")

	api.Get("/", handlers.GetAllJobs)
}
