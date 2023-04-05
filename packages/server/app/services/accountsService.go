package services

import (
	"errors"

	models "go.labs/server/app/models"
)

type accountsService struct {
	model *models.AccountModel
}

func (accountService *accountsService) GetAllAccounts() []models.Account {
	return accountService.model.FindAll()
}

func (accountService *accountsService) GetOneAccount(id int) *models.Account {
	return accountService.model.FindOne(id)
}

func (accountService *accountsService) GetOneByEmail(email string) *models.Account {
	return accountService.model.FindOneByEmail(email)
}

func (accountService *accountsService) AddAccount(account *models.Account) error {
	if accountService.model.ExistsByEmail(account.Email) {
		return errors.New("account with this email already exists")
	}

	accountService.model.Add(account)
	return nil
}

func (accountService *accountsService) RemoveAccount(id int) {
	accountService.model.Delete(id)
}
