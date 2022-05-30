package main

import (
	"webscraper/server"
	"webscraper/server/models"
)

func main() {
	models.Setup()
	server.SetupAndListen()
}
