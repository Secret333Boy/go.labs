package accounts

import (
	"encoding/json"
	"net/http"

	"go.labs/server/app/middlewares"
	"go.labs/server/app/router"
)

func GetAccountsRouter() *router.Router {
	router := router.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		account, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		err = json.NewEncoder(w).Encode(account)
		if err != nil {
			http.Error(w, "Failed reading json file", http.StatusInternalServerError)
		}
	})
	return router
}
