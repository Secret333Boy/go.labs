package middlewares

import "net/http"

func UseJSONContentType(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}
