package main

import (
	"log"

	"github.com/joho/godotenv"
	"go.labs/server/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Run()
}
