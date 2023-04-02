package services

import (
	models "go.labs/server/app/models"
)

type AccountsService struct {
	model *models.AccountModel
}

func (accountService *AccountsService) GetAllAccounts() []models.Account {
	return accountService.model.FindAll()
}

func (accountService *AccountsService) GetOneAccount(id int) *models.Account {
	return accountService.model.FindOne(id)
}

func (accountService *AccountsService) AddAccount(account *models.Account) {
	accountService.model.Add(account)
}

func (accountService *AccountsService) RemoveAccount(id int) {
	accountService.model.Delete(id)
}
