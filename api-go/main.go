package main

import (
	"log"

	"api-go/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	}

	router.CreateRouter()
}
