package main

import (
	"book-api/config"
	"book-api/routes"
)

func main() {
	config.ConnectDatabase()

	router := routes.SetupRouter()
	router.Run(":8080")
}
