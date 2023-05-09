package account

import (
	"errors"
	"fmt"

	"go.labs/server/app/models"
	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(DB *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: DB}
}

func (ar *AccountRepository) FindAllAccounts() []models.Account {
	var accounts []models.Account

	if result := ar.DB.Find(&accounts); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return accounts
}

func (ar *AccountRepository) FindOneAccount(id int) *models.Account {
	account := &models.Account{}

	if result := ar.DB.First(account, id); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return account
}

func (ar *AccountRepository) FindOneAccountByEmail(email string) *models.Account {
	account := &models.Account{}

	if result := ar.DB.Where("email = ?", email).First(account); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return account
}

func (ar *AccountRepository) CreateAccount(account *models.Account) error {
	if result := ar.DB.Create(account); result.Error != nil {
		fmt.Println(result.Error)
		return errors.New("error while saving account")
	}

	return nil
}

func (ar *AccountRepository) DeleteAccount(account *models.Account) {
	if result := ar.DB.Delete(account); result.Error != nil {
		fmt.Println(result.Error)
	}
}
