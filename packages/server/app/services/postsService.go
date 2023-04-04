package services

import (
	"go.labs/server/app/controllers/api/posts/dtos"
	"go.labs/server/app/models"
	"time"
)

type postsService struct {
	model *models.PostModel
}

func (postsService *postsService) GetAllPosts(limit int, offset int) []models.Post {
	return postsService.model.FindAll(limit, offset)
}

func (postsService *postsService) GetOnePost(id int) *models.Post {
	return postsService.model.FindOne(id)
}

//func (postsService *postsService) GetOneByEmail(email string) *models.Post {
//	return postsService.model.FindOneByEmail(email)
//}

func (postsService *postsService) AddPost(account *models.Account, createPostDto *dtos.CreatePostDto) {
	post := &models.Post{Account: account, Title: createPostDto.Title, Description: createPostDto.Description, PublishedAt: time.Now()}
	postsService.model.Add(post)
}

func (postsService *postsService) RemovePost(id int) {
	postsService.model.Delete(id)
}
