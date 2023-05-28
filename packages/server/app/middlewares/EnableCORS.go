package middlewares

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func EnableCors(next httprouter.Handle) httprouter.Handle {
	clientURL, clientURLExists := os.LookupEnv("CLIENT_URL")

	if !clientURLExists {
		clientURL = ""
	}

	return func(w http.ResponseWriter,
		r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", clientURL)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next(w, r, ps)
	}
}
