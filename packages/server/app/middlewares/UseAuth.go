package middlewares

import (
	"net/http"
	"strings"

	"go.labs/server/app/models"
	"go.labs/server/app/services/auth"
)

func UseAuth(w http.ResponseWriter, r *http.Request) (*models.Account, error) {
	header := r.Header.Get("Authorization")

	token := strings.Replace(header, "Bearer ", "", 1)
	account, err := auth.GetAuthServiceInstance().Validate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return nil, err
	}

	return account, err
}
