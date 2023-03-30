package services

import (
	"go.labs/server/app/models"
)

func NewAccountsService() *AccountsService {
	var accountModel = models.NewAccountModel()
	return &AccountsService{accountModel}
}
