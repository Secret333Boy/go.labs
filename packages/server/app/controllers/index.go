package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.labs/server/app/controllers/api/accounts"
	"go.labs/server/app/controllers/api/auth"
	"go.labs/server/app/controllers/api/posts"
	"go.labs/server/app/db"
	"go.labs/server/app/middlewares"
	accountsService "go.labs/server/app/services/accounts"
	authService "go.labs/server/app/services/auth"
	postsService "go.labs/server/app/services/posts"
)

func GetIndexRouter() *httprouter.Router {
	db := db.Init()

	router := httprouter.New()

	router.GET("/api", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "go.labs v1.0")
	})

	accountsService := &accountsService.AccountsService{DB: db}
	authService := &authService.AuthService{DB: db, AccountsService: accountsService}
	postsService := &postsService.PostsService{DB: db}

	useAuthMiddleware := middlewares.NewUseAuthMiddleware(authService)

	authHandler := auth.NewAuthHandler(authService)
	accountsHandler := accounts.NewAccountsHandler(accountsService, useAuthMiddleware)
	postsHandler := posts.NewPostsHandler(postsService, useAuthMiddleware)

	authHandler.RegisterHandler(router)
	accountsHandler.RegisterHandler(router)
	postsHandler.RegisterHandler(router)

	return router
}
