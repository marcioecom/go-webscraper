package main

import (
	"log"
	"webscraper/server"
	"webscraper/server/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error to load dot env")
	}

	models.Setup()
	server.SetupAndListen()
}
