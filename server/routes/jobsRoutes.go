package routes

import (
	"webscraper/server/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupJobsRoutes(app *fiber.Router) {
	api := (*app).Group("/jobs")

	api.Get("/", handler.GetAllJobs)
}
