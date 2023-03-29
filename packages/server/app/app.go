package app

import (
	"fmt"
	"log"
	"net/http"

	accountsController "go.labs/server/app/controllers/accounts"
)

func Run() {
	router := Router{}

	// router.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "go.labs API 1.0")
	// })

	router.Handle("/accounts", accountsController.GetAccountsRouter())

	fmt.Println("Server started on port " + "8081")
	log.Fatal(http.ListenAndServe(":8081", &router))
}
