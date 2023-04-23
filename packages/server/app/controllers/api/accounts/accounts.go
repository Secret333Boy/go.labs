package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.labs/server/app/middlewares"
	"go.labs/server/app/services/accounts"
)

type AccountsHandler struct {
	service           *accounts.AccountsService
	useAuthMiddleware *middlewares.UseAuthMiddleware
}

func NewAccountsHandler(service *accounts.AccountsService, useAuthMiddleware *middlewares.UseAuthMiddleware) *AccountsHandler {
	return &AccountsHandler{service: service, useAuthMiddleware: useAuthMiddleware}
}

func (h *AccountsHandler) RegisterHandler(router *httprouter.Router) {
	router.GET("/api/accounts", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		account, err := h.useAuthMiddleware.UseAuth(w, r)
		if err != nil {
			return
		}

		err = json.NewEncoder(w).Encode(account)
		if err != nil {
			http.Error(w, "Failed reading json file", http.StatusInternalServerError)
		}
	})
}
