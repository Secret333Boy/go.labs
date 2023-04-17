package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.labs/server/app/middlewares"
)

func HandleAccounts(router *httprouter.Router) {
	router.GET("/api/accounts", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		account, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		err = json.NewEncoder(w).Encode(account)
		if err != nil {
			http.Error(w, "Failed reading json file", http.StatusInternalServerError)
		}
	})
}
