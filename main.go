package main

import (
	"webscraper/server"
	"webscraper/server/models"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./.env")

	models.Setup()
	server.SetupAndListen()
}
