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
	accountRepository "go.labs/server/app/repositories/account"
	messageRepository "go.labs/server/app/repositories/message"
	postRepository "go.labs/server/app/repositories/post"
	tokenRepository "go.labs/server/app/repositories/token"
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

	accountRepository := accountRepository.NewAccountRepository(db)
	postRepository := postRepository.NewPostRepository(db)
	messageRepository := messageRepository.NewMessageRepository(db)
	tokenRepository := tokenRepository.NewTokenRepository(db)

	accountsService := accountsService.NewAccountsService(accountRepository)
	authService := authService.NewAuthService(tokenRepository, accountsService)
	postsService := postsService.NewPostsService(postRepository, messageRepository)

	useAuthMiddleware := middlewares.NewUseAuthMiddleware(authService)

	authHandler := auth.NewAuthHandler(authService)
	accountsHandler := accounts.NewAccountsHandler(accountsService, useAuthMiddleware)
	postsHandler := posts.NewPostsHandler(postsService, useAuthMiddleware)

	authHandler.RegisterHandler(router)
	accountsHandler.RegisterHandler(router)
	postsHandler.RegisterHandler(router)

	return router
}
