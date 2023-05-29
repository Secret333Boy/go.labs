package posts

import (
	"go.labs/server/app/models"
)

type PostRepositoryMock struct {
	Post        *models.Post
	GetError    error
	CreateError error
	DeleteError error
}

func (pr *PostRepositoryMock) CreatePost(*models.Account, string, string) error {
	return pr.CreateError
}

func (pr *PostRepositoryMock) RemovePost(*models.Post, *models.Account) error {
	return pr.DeleteError
}

func (pr *PostRepositoryMock) FindOnePost(int) (*models.Post, error) {
	return pr.Post, pr.GetError
}
func (pr *PostRepositoryMock) UpdatePost(*models.Post, *models.Account, string, string) error {
	//TODO implement me
	panic("implement me")
}
func (pr *PostRepositoryMock) FindAllPosts(int, int) []models.Post {
	//TODO implement me
	panic("implement me")
}
