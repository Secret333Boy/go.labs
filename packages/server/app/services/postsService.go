package services

import (
	"go.labs/server/app/controllers/api/posts/dtos"
	"go.labs/server/app/models"
	"time"
)

type postsService struct {
	model *models.PostModel
}

func (postService *postsService) GetAllPosts(limit int, offset int) []models.Post {
	return postService.model.FindAll(limit, offset)
}

func (postService *postsService) GetOnePost(id int) *models.Post {
	return postService.model.FindOne(id)
}

//func (postService *postsService) GetOneByEmail(email string) *models.Post {
//	return postService.model.FindOneByEmail(email)
//}

func (postService *postsService) AddPost(account *models.Account, createPostDto *dtos.CreatePostDto) {
	post := &models.Post{Account: account, Title: createPostDto.Title, Description: createPostDto.Description, PublishedAt: time.Now()}
	postService.model.Add(post)
}

func (postService *postsService) RemovePost(id int) {
	postService.model.Delete(id)
}
