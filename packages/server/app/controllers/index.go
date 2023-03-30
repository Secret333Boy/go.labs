package controllers

import (
	"go.labs/server/app/controllers/api"
	"go.labs/server/app/router"
)

func GetIndexRouter() *router.Router {
	router := router.NewRouter()

	router.Use("/api", api.GetAPIRouter())

	return router
}
