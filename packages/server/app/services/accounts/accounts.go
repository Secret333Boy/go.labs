package accounts

import (
	"errors"

	models "go.labs/server/app/models"
)

type accountsRepository interface {
	FindAllAccounts() []models.Account
	FindOneAccount(id int) *models.Account
	FindOneAccountByEmail(email string) *models.Account
	CreateAccount(account *models.Account) error
	DeleteAccount(account *models.Account)
}

type AccountsService struct {
	accountsRepository accountsRepository
}

func NewAccountsService(accountsRepository accountsRepository) Account {
	return &AccountsService{accountsRepository: accountsRepository}
}

type Account interface {
	GetAllAccounts() []models.Account
	GetOneByEmail(email string) *models.Account
	AddAccount(account *models.Account) error
	RemoveAccount(id int)
}

func (a *AccountsService) GetAllAccounts() []models.Account {
	return a.accountsRepository.FindAllAccounts()
}

func (a *AccountsService) GetOneAccount(id int) *models.Account {
	return a.accountsRepository.FindOneAccount(id)
}

func (a *AccountsService) GetOneByEmail(email string) *models.Account {
	return a.accountsRepository.FindOneAccountByEmail(email)
}

func (a *AccountsService) AddAccount(account *models.Account) error {
	if a.GetOneByEmail(account.Email) != nil {
		return errors.New("account with this email already exists")
	}

	return a.accountsRepository.CreateAccount(account)
}

func (a *AccountsService) RemoveAccount(id int) {
	account := a.GetOneAccount(id)
	a.accountsRepository.DeleteAccount(account)
}
