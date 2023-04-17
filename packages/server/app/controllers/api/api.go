package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.labs/server/app/controllers/api/accounts"
	"go.labs/server/app/controllers/api/auth"
	"go.labs/server/app/controllers/api/posts"
)

func HandleAPI(router *httprouter.Router) {
	router.GET("/api", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "go.labs v1.0")
	})

	accounts.HandleAccounts(router)
	auth.HandleAuth(router)
	posts.HandlePosts(router)
}
