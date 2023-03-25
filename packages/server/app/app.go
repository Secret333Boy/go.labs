package app

import (
	"fmt"
	"log"
	"net/http"

	accountsController "go.labs/app/controllers/accounts"
)

func Run() {
	router := http.NewServeMux()

	router.Handle("/accounts", accountsController.GetAccountsRouter())

	fmt.Println("Server started on port " + "8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
