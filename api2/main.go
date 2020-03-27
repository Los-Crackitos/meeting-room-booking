package main

import (
	"log"

	"api2/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	}

	router.CreateRouter()

}
