package accounts

import (
	"errors"
	"fmt"

	models "go.labs/server/app/models"
	"gorm.io/gorm"
)

type AccountsService struct {
	DB *gorm.DB
}

func (a *AccountsService) GetAllAccounts() []models.Account {
	var accounts []models.Account

	if result := a.DB.Find(&accounts); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return accounts
}

func (a *AccountsService) GetOneAccount(id int) *models.Account {
	account := &models.Account{}

	if result := a.DB.First(account, id); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return account
}

func (a *AccountsService) GetOneByEmail(email string) *models.Account {
	account := &models.Account{}

	if result := a.DB.Where("email = ?", email).First(account); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return account
}

func (a *AccountsService) AddAccount(account *models.Account) error {
	if a.GetOneByEmail(account.Email) != nil {
		return errors.New("account with this email already exists")
	}

	if result := a.DB.Create(account); result.Error != nil {
		fmt.Println(result.Error)
		return errors.New("error while saving account")
	}

	return nil
}

func (a *AccountsService) RemoveAccount(id int) {
	account := a.GetOneAccount(id)

	if result := a.DB.Delete(account); result.Error != nil {
		fmt.Println(result.Error)
	}
}
