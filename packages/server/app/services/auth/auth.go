package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"go.labs/server/app/models"
	"go.labs/server/app/services/accounts"
	"go.labs/server/app/services/tokens"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	tokenModel      *models.TokenModel
	accountsService *accounts.AccountsService
}

func (a *AuthService) Register(email string, password string) (*tokens.Tokens, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	hash := string(bytes)

	accessToken, err := tokens.Encrypt(email, hash, false)
	if err != nil {
		return nil, err
	}

	refreshToken, err := tokens.Encrypt(email, hash, true)
	if err != nil {
		return nil, err
	}

	account := &models.Account{Email: email, Hash: hash}
	a.tokenModel.Add(&models.Token{Account: account, RefreshToken: refreshToken})

	err = a.accountsService.AddAccount(account)
	if err != nil {
		return nil, err
	}

	return &tokens.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (a *AuthService) Login(email string, password string) (*tokens.Tokens, error) {
	account := a.accountsService.GetOneByEmail(email)
	if account == nil {
		return nil, errors.New("account not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Hash), []byte(password))
	if err != nil {
		return nil, err
	}

	accessToken, err := tokens.Encrypt(email, account.Hash, false)
	if err != nil {
		return nil, err
	}

	refreshToken, err := tokens.Encrypt(email, account.Hash, true)
	if err != nil {
		return nil, err
	}

	a.tokenModel.UpdateByAccount(account, refreshToken)

	return &tokens.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (a *AuthService) Validate(tokenString string) (*models.Account, error) {
	token, err := tokens.Decrypt(tokenString, false)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid JWT Token")
	}

	email := fmt.Sprint(claims["email"])
	hash := fmt.Sprint(claims["hash"])

	account := a.accountsService.GetOneByEmail(email)

	if account == nil {
		return nil, errors.New("account not found")
	}

	if account.Hash != hash {
		return nil, errors.New("hash is invalid")
	}

	return account, nil
}

func (a *AuthService) Refresh(tokenString string) (*tokens.Tokens, error) {
	token, err := tokens.Decrypt(tokenString, true)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid JWT Token")
	}

	email := fmt.Sprint(claims["email"])
	hash := fmt.Sprint(claims["hash"])

	account := a.accountsService.GetOneByEmail(email)

	if account == nil {
		return nil, errors.New("account not found")
	}

	if account.Hash != hash {
		return nil, errors.New("hash is invalid")
	}

	accessToken, err := tokens.Encrypt(account.Email, account.Hash, false)
	if err != nil {
		return nil, err
	}

	refreshToken, err := tokens.Encrypt(account.Email, account.Hash, true)
	if err != nil {
		return nil, err
	}

	a.tokenModel.UpdateByAccount(account, refreshToken)

	return &tokens.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

var authServiceInstance = &AuthService{models.NewTokenModel(), accounts.GetAccountsServiceInstance()}

func GetAuthServiceInstance() *AuthService {
	return authServiceInstance
}
