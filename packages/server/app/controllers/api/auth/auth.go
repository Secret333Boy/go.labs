package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"go.labs/server/app/controllers/api/auth/dtos"
	"go.labs/server/app/router"
	"go.labs/server/app/services"
)

func GetAuthRouter() *router.Router {
	router := router.NewRouter()
	authService := services.AuthService

	router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		registerDto := &dtos.RegisterDto{}

		err := json.NewDecoder(r.Body).Decode(registerDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validationErr := registerDto.Validate()
		if validationErr != nil {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}

		tokens, err := authService.Register(registerDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: tokens.RefreshToken, HttpOnly: true, Path: "/api", MaxAge: services.RefreshTokenExpirationSeconds})
		err = json.NewEncoder(w).Encode(tokens.AccessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		loginDto := &dtos.RegisterDto{}

		err := json.NewDecoder(r.Body).Decode(loginDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validationErr := loginDto.Validate()
		if validationErr != nil {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}

		tokens, err := authService.Login(loginDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: tokens.RefreshToken, HttpOnly: true, Path: "/api", MaxAge: services.RefreshTokenExpirationSeconds})
		err = json.NewEncoder(w).Encode(tokens.AccessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	router.Get("/validate", func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		token := strings.Replace(header, "Bearer ", "", 1)
		err := authService.Validate(token)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
	})

	router.Get("/refresh", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("refreshToken")

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		tokens, err := authService.Refresh(tokenString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		w.Header().Add("Content-Type", "application/json")
		http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: tokens.RefreshToken, HttpOnly: true, Path: "/api", MaxAge: services.RefreshTokenExpirationSeconds})
		err = json.NewEncoder(w).Encode(tokens.AccessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	router.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: "", HttpOnly: true, Path: "/api", MaxAge: -1})
	})

	return router
}
