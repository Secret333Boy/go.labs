package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"go.labs/server/app/middlewares"
	"go.labs/server/app/services/auth"
	tokensService "go.labs/server/app/services/tokens"
)

func HandleAuth(router *httprouter.Router) {
	authService := auth.GetAuthServiceInstance()

	router.POST("/api/auth/register", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		registerDto := &RegisterDto{}

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

		tokens, err := authService.Register(registerDto.Email, registerDto.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		middlewares.UseJSONContentType(w)
		http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: tokens.RefreshToken, HttpOnly: true, Path: "/api", MaxAge: tokensService.RefreshTokenExpirationSeconds})
		err = json.NewEncoder(w).Encode(tokens.AccessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	router.POST("/api/auth/login", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		loginDto := &RegisterDto{}

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

		tokens, err := authService.Login(loginDto.Email, loginDto.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		middlewares.UseJSONContentType(w)
		http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: tokens.RefreshToken, HttpOnly: true, Path: "/api", MaxAge: tokensService.RefreshTokenExpirationSeconds})
		err = json.NewEncoder(w).Encode(tokens.AccessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	router.GET("/api/auth/validate", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		header := r.Header.Get("Authorization")

		token := strings.Replace(header, "Bearer ", "", 1)
		_, err := authService.Validate(token)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
	})

	router.GET("/api/auth/refresh", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

		middlewares.UseJSONContentType(w)
		http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: tokens.RefreshToken, HttpOnly: true, Path: "/api", MaxAge: tokensService.RefreshTokenExpirationSeconds})
		err = json.NewEncoder(w).Encode(tokens.AccessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	router.GET("/api/auth/logout", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: "", HttpOnly: true, Path: "/api", MaxAge: -1})
	})

}
