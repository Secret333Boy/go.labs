package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.labs/server/app/controllers"
)

const DefaultPort = "8081"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, portExists := os.LookupEnv("PORT")

	if !portExists {
		port = DefaultPort
	}

	router := controllers.GetIndexRouter()

	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
