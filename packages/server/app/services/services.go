package services

import (
	"go.labs/server/app/models"
)

var AccountsService = &accountsService{models.NewAccountModel()}
var AuthService = &authService{models.NewTokenModel(), AccountsService}
var PostsService = &postsService{models.NewPostModel()}
