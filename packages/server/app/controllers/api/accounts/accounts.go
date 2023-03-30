package accounts

import (
	"encoding/json"
	"net/http"

	"go.labs/server/app/models"
	"go.labs/server/app/router"
	"go.labs/server/app/services"
)

func GetAccountsRouter() *router.Router {
	router := router.NewRouter()
	accountsService := services.NewAccountsService()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(accountsService.GetAllAccounts())
		if err != nil {
			http.Error(w, "Failed reading json file", http.StatusInternalServerError)
		}
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var account models.Account
		json.NewDecoder(r.Body).Decode(&account)
		accountsService.AddAccount(&account)
	})

	return router
}
