package services

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"go.labs/server/app/controllers/api/auth/dtos"
	"go.labs/server/app/models"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	tokenModel      *models.TokenModel
	accountsService *accountsService
}

func (authService *authService) Register(registerDto *dtos.RegisterDto) (*Tokens, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), 14)
	if err != nil {
		return nil, err
	}

	hash := string(bytes)

	accessToken, err := Encrypt(registerDto.Email, hash, false)
	if err != nil {
		return nil, err
	}

	refreshToken, err := Encrypt(registerDto.Email, hash, true)
	if err != nil {
		return nil, err
	}

	account := &models.Account{Email: registerDto.Email, Hash: hash}
	authService.tokenModel.Add(&models.Token{Account: account, RefreshToken: refreshToken})

	err = authService.accountsService.AddAccount(account)
	if err != nil {
		return nil, err
	}

	return &Tokens{accessToken, refreshToken}, nil
}

func (authService *authService) Login(loginDto *dtos.RegisterDto) (*Tokens, error) {
	account := authService.accountsService.GetOneByEmail(loginDto.Email)
	if account == nil {
		return nil, errors.New("account not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Hash), []byte(loginDto.Password))
	if err != nil {
		return nil, err
	}

	accessToken, err := Encrypt(loginDto.Email, account.Hash, false)
	if err != nil {
		return nil, err
	}

	refreshToken, err := Encrypt(loginDto.Email, account.Hash, true)
	if err != nil {
		return nil, err
	}

	authService.tokenModel.UpdateByAccount(account, refreshToken)

	return &Tokens{accessToken, refreshToken}, nil
}

func (authService *authService) Validate(tokenString string) error {
	token, err := Decrypt(tokenString, false)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return errors.New("invalid JWT Token")
	}

	email := fmt.Sprint(claims["email"])
	hash := fmt.Sprint(claims["hash"])

	account := authService.accountsService.GetOneByEmail(email)

	if account == nil {
		return errors.New("account not found")
	}

	if account.Hash != hash {
		return errors.New("hash is invalid")
	}

	return nil
}

func (authService *authService) Refresh(tokenString string) (*Tokens, error) {
	token, err := Decrypt(tokenString, true)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid JWT Token")
	}

	email := fmt.Sprint(claims["email"])
	hash := fmt.Sprint(claims["hash"])

	account := authService.accountsService.GetOneByEmail(email)

	if account == nil {
		return nil, errors.New("account not found")
	}

	if account.Hash != hash {
		return nil, errors.New("hash is invalid")
	}

	accessToken, err := Encrypt(account.Email, account.Hash, false)
	if err != nil {
		return nil, err
	}

	refreshToken, err := Encrypt(account.Email, account.Hash, true)
	if err != nil {
		return nil, err
	}

	authService.tokenModel.UpdateByAccount(account, refreshToken)

	return &Tokens{accessToken, refreshToken}, nil
}
