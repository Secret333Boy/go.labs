package app

import (
	"fmt"
	"log"
	"net/http"

	"go.labs/server/app/controllers"
)

func Run() {

	var router = controllers.GetIndexRouter()

	fmt.Println("Server started on port " + "8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
