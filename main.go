package main

import (
	"webscraper/server"
	"webscraper/server/crawler"
	"webscraper/server/models"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./.env")

	models.Setup()
	crawler.SchedulerJob()
	server.SetupAndListen()
}
