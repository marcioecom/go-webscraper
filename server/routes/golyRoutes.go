package routes

import (
	"webscraper/server/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupGolyRoutes(app *fiber.Router) {
	api := (*app).Group("/goly")

	api.Get("/:redirect", handler.Redirect)
	api.Post("/", handler.CreateGoly)
}
