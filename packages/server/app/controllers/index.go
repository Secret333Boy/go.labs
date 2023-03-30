package controllers

import (
	"fmt"
	"net/http"

	"go.labs/server/app/controllers/api/accounts"
	"go.labs/server/app/router"
)

func GetIndexRouter() *router.Router {
	router := router.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "go.labs v1.0")
	})

	router.Use("/accounts", accounts.GetAccountsRouter())

	return router
}
