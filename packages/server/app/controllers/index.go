package controllers

import (
	"github.com/julienschmidt/httprouter"
	"go.labs/server/app/controllers/api"
)

func GetIndexRouter() *httprouter.Router {
	router := httprouter.New()

	api.HandleAPI(router)

	return router
}
