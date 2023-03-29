package accountsController

import (
	"encoding/json"
	"net/http"

	models "go.labs/server/app/models"
	accountsService "go.labs/server/app/services/accounts"
)

func GetAccountsRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", indexHandler)
	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var err = json.NewEncoder(w).Encode(accountsService.GetAllAccounts())
		if err != nil {
			http.Error(w, "Failed read json file", http.StatusInternalServerError)
		}

	case http.MethodPost:
		var account models.Account
		json.NewDecoder(r.Body).Decode(&account)
		accountsService.AddAccount(&account)
	default:
		http.Error(w, "Method is not accepted", http.StatusMethodNotAllowed)
	}
}
