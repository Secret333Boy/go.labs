package models

type Token struct {
	Id           int
	Account      *Account
	RefreshToken string
	model        *TokenModel
}

type TokenModel struct {
	tokens []Token
	lastId int
}

func NewTokenModel() *TokenModel {
	model := new(TokenModel)
	model.tokens = make([]Token, 0)
	return model
}

func (model *TokenModel) FindAll() []Token {
	return model.tokens
}

func (model *TokenModel) FindOne(id int) *Token {
	for _, token := range model.tokens {
		if token.Id == id {
			return &token
		}
	}

	return nil
}

func (model *TokenModel) FindOneByAccount(account *Account) *Token {
	for _, token := range model.tokens {
		if token.Account == account {
			return &token
		}
	}

	return nil
}

func (model *TokenModel) Add(token *Token) {
	token.model = model
	token.Id = model.lastId + 1
	model.tokens = append(model.tokens, *token)
	model.lastId++
}

func (model *TokenModel) Delete(id int) {
	for i, token := range model.tokens {
		if token.Id == id {
			token.model = nil
			model.tokens = append(model.tokens[:i], model.tokens[i+1:]...)
			return
		}
	}
}

func (model *TokenModel) Exists(id int) bool {
	for _, token := range model.tokens {
		if token.Id == id {
			return true
		}
	}
	return false
}

func (model *TokenModel) UpdateByAccount(account *Account, refreshToken string) *Token {
	token := model.FindOneByAccount(account)
	if token == nil {
		return nil
	}

	token.RefreshToken = refreshToken

	return token
}
