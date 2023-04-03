package accounts

import (
	"encoding/json"
	"net/http"

	"go.labs/server/app/router"
	"go.labs/server/app/services"
)

func GetAccountsRouter() *router.Router {
	router := router.NewRouter()
	accountsService := services.AccountsService

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(accountsService.GetAllAccounts())
		if err != nil {
			http.Error(w, "Failed reading json file", http.StatusInternalServerError)
		}
	})
	return router
}
