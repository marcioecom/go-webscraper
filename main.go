package main

import (
	"webscraper/server"
	"webscraper/server/crawler"
	"webscraper/server/models"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./.env")

	crawler.SchedulerJob()
	models.Setup()
	server.SetupAndListen()
}
