package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"go.labs/server/app/models"
	"go.labs/server/app/services/accounts"
	"go.labs/server/app/services/tokens"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB      *gorm.DB
	Account accounts.Account
}

func NewAuthService(db *gorm.DB, account accounts.Account) Auth {
	return &AuthService{DB: db, Account: account}
}

type Auth interface {
	Register(email string, password string) (*tokens.Tokens, error)
	Login(email string, password string) (*tokens.Tokens, error)
	Validate(tokenString string) (*models.Account, error)
	Refresh(tokenString string) (*tokens.Tokens, error)
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

	err = a.Account.AddAccount(account)
	if err != nil {
		return nil, err
	}

	a.DB.Create(&models.Token{Account: *account, RefreshToken: refreshToken})

	return &tokens.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (a *AuthService) Login(email string, password string) (*tokens.Tokens, error) {
	account := a.Account.GetOneByEmail(email)
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

	a.DB.Model(&models.Token{}).Where("account_id = ?", account.ID).Updates(&models.Token{RefreshToken: refreshToken})

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

	account := a.Account.GetOneByEmail(email)

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

	account := a.Account.GetOneByEmail(email)

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

	if result := a.DB.Model(&models.Token{}).Where("account_id = ?", account.ID).Update("refresh_token", refreshToken); result.Error != nil {
		fmt.Println(result.Error)
		return nil, errors.New("failed to save refresh token")
	}

	return &tokens.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
