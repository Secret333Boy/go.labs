package api

import (
	"fmt"
	"go.labs/server/app/controllers/api/posts"
	"net/http"

	"go.labs/server/app/controllers/api/accounts"
	"go.labs/server/app/controllers/api/auth"
	"go.labs/server/app/router"
)

func GetAPIRouter() *router.Router {
	router := router.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "go.labs v1.0")
	})

	router.Use("/accounts", accounts.GetAccountsRouter())
	router.Use("/auth", auth.GetAuthRouter())
	router.Use("/posts", posts.GetPostsRouter())

	return router
}
