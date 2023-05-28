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

type AuthHandler struct {
	service auth.Auth
}

func NewAuthHandler(service auth.Auth) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterHandler(router *httprouter.Router) {
	router.POST("/api/auth/register", middlewares.EnableCors(h.register))

	router.POST("/api/auth/login", middlewares.EnableCors(h.login))

	router.GET("/api/auth/validate", middlewares.EnableCors(h.validate))

	router.GET("/api/auth/refresh", middlewares.EnableCors(h.refresh))

	router.GET("/api/auth/logout", middlewares.EnableCors(h.logout))
}

func (h *AuthHandler) register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	tokens, err := h.service.Register(registerDto.Email, registerDto.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	middlewares.UseJSONContentType(w)
	http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: tokens.RefreshToken, HttpOnly: true, Path: "/api", MaxAge: tokensService.RefreshTokenExpirationSeconds, Secure: true, SameSite: http.SameSiteNoneMode})
	err = json.NewEncoder(w).Encode(tokens.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	tokens, err := h.service.Login(loginDto.Email, loginDto.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	middlewares.UseJSONContentType(w)
	http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: tokens.RefreshToken, HttpOnly: true, Path: "/api", MaxAge: tokensService.RefreshTokenExpirationSeconds, Secure: true, SameSite: http.SameSiteNoneMode})
	err = json.NewEncoder(w).Encode(tokens.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *AuthHandler) validate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := r.Header.Get("Authorization")

	token := strings.Replace(header, "Bearer ", "", 1)
	_, err := h.service.Validate(token)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
}

func (h *AuthHandler) refresh(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, err := r.Cookie("refreshToken")

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenString := cookie.Value

	tokens, err := h.service.Refresh(tokenString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	middlewares.UseJSONContentType(w)
	http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: tokens.RefreshToken, HttpOnly: true, Path: "/api", MaxAge: tokensService.RefreshTokenExpirationSeconds, Secure: true, SameSite: http.SameSiteNoneMode})
	err = json.NewEncoder(w).Encode(tokens.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *AuthHandler) logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.SetCookie(w, &http.Cookie{Name: "refreshToken", Value: "", HttpOnly: true, Path: "/api", MaxAge: -1})
}
