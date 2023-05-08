package token

import (
	"errors"
	"fmt"

	"go.labs/server/app/models"
	"gorm.io/gorm"
)

type TokenRepository struct {
	DB *gorm.DB
}

func NewTokenRepository(DB *gorm.DB) *TokenRepository {
	return &TokenRepository{DB: DB}
}

func (tr *TokenRepository) CreateToken(token *models.Token) {
	tr.DB.Create(token)
}

func (tr *TokenRepository) UpdateToken(account *models.Account, refreshToken string) error {
	if result := tr.DB.Model(&models.Token{}).Where("account_id = ?", account.ID).Update("refresh_token", refreshToken); result.Error != nil {
		fmt.Println(result.Error)
		return errors.New("failed to save refresh token")
	}

	return nil
}
