package AccountModel

type Account struct {
	Id        int
	FirstName string
	LastName  string
	model     *AccountModel
}

type AccountModel struct {
	accounts []Account
}

func NewAccountModel() *AccountModel {
	var model = new(AccountModel)
	model.accounts = make([]Account, 0)
	return model
}

func (model *AccountModel) FindAll() []Account {
	return model.accounts
}

func (model *AccountModel) FindOne(id int) *Account {
	for _, account := range model.accounts {
		if account.Id == id {
			return &account
		}
	}

	return nil
}

func (model *AccountModel) Add(account *Account) {
	account.model = model
	model.accounts = append(model.accounts, *account)
}

func (model *AccountModel) Delete(id int) {
	for i, account := range model.accounts {
		if account.Id == id {
			account.model = nil
			model.accounts = append(model.accounts[:i], model.accounts[i+1:]...)
			return
		}
	}
}

func (model *AccountModel) Exists(id int) bool {
	for _, account := range model.accounts {
		if account.Id == id {
			return true
		}
	}
	return false
}
