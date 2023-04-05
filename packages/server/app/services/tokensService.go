package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

var accessTokenSecret = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
var refreshTokenSecret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

const AccessTokenExpirationSeconds = 10_800     // 3 hours
const RefreshTokenExpirationSeconds = 2_592_000 // 30 days

func Encrypt(email string, hash string, isRefreshToken bool) (string, error) {
	secret := accessTokenSecret
	if isRefreshToken {
		secret = refreshTokenSecret
	}

	expirationTime := time.Now().Add(AccessTokenExpirationSeconds * time.Second)
	if isRefreshToken {
		expirationTime = time.Now().Add(RefreshTokenExpirationSeconds * time.Second)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"hash":  hash,
		"exp":   expirationTime.Unix(),
	})

	return token.SignedString(secret)
}

func Decrypt(tokenString string, isRefreshToken bool) (*jwt.Token, error) {
	secret := accessTokenSecret
	if isRefreshToken {
		secret = refreshTokenSecret
	}

	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return secret, nil
	})
}
