package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go.labs/server/app/controllers"
)

const DefaultPort = "8081"

func Run() {
	port, portExists := os.LookupEnv("PORT")

	if !portExists {
		port = DefaultPort
	}

	router := controllers.GetIndexRouter()

	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
