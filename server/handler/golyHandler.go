package handler

import (
	"fmt"
	"webscraper/server/models"
	"webscraper/utils"

	"github.com/gofiber/fiber/v2"
)

func Redirect(c *fiber.Ctx) error {
	golyUrl := c.Params("redirect")

	goly, err := models.FindByGolyUrl(golyUrl)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "could not find goly in db " + err.Error(),
		})
	}

	goly.Clicked += 1
	err = models.UpdateGoly(goly)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

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
