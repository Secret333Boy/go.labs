package accounts

import (
	"errors"

	models "go.labs/server/app/models"
)

type AccountsService struct {
	model *models.AccountModel
}

func (a *AccountsService) GetAllAccounts() []models.Account {
	return a.model.FindAll()
}

func (a *AccountsService) GetOneAccount(id int) *models.Account {
	return a.model.FindOne(id)
}

func (a *AccountsService) GetOneByEmail(email string) *models.Account {
	return a.model.FindOneByEmail(email)
}

func (a *AccountsService) AddAccount(account *models.Account) error {
	if a.model.ExistsByEmail(account.Email) {
		return errors.New("account with this email already exists")
	}

	a.model.Add(account)
	return nil
}

func (a *AccountsService) RemoveAccount(id int) {
	a.model.Delete(id)
}

var accountsServiceInstance = &AccountsService{models.NewAccountModel()}

func GetAccountsServiceInstance() *AccountsService {
	return accountsServiceInstance
}
