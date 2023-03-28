package accountsService

import (
	models "go.labs/server/app/models"
)

var accountModel = models.NewAccountModel()

func GetAllAccounts() []models.Account {
	return accountModel.FindAll()
}

func GetOneAccount(id int) *models.Account {
	return accountModel.FindOne(id)
}

func AddAccount(account *models.Account) {
	accountModel.Add(account)
}

func RemoveAccount(id int) {
	accountModel.Delete(id)
}
