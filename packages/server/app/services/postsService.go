package services

import (
	"errors"
	"go.labs/server/app/controllers/api/posts/dtos"
	"go.labs/server/app/models"
	"time"
)

type postsService struct {
	postModel    *models.PostModel
	messageModel *models.MessageModel
}

func (postsService *postsService) GetAllPosts(limit int, offset int) []models.Post {
	return postsService.postModel.FindAll(limit, offset)
}

func (postsService *postsService) GetOnePost(id int) *models.Post {
	return postsService.postModel.FindOne(id)
}

func (postsService *postsService) AddPost(account *models.Account, createPostDto *dtos.CreatePostDto) {
	post := &models.Post{Account: account, Title: createPostDto.Title, Description: createPostDto.Description, PublishedAt: time.Now()}
	postsService.postModel.Add(post)
}

func (postsService *postsService) RemovePost(account *models.Account, id int) error {
	if !postsService.postModel.Exists(id) {
		return errors.New("post not found")
	}
	if postsService.postModel.FindAccountByPostId(id) != account {
		return errors.New("access denied")
	}
	postsService.postModel.Delete(id)
	return nil
}

//
//func (postsService *postsService) GetAllMessagesByPostId(limit int, offset int) []models.Post {
//	return postsService.messageModel.FindAll(limit, offset)
//}
//
//func (postsService *postsService) GetMessageByPostId(id int) *models.Post {
//	return postsService.messageModel.FindOne(id)
//}
//
////func (postsService *postsService) GetOneByEmail(email string) *models.Post {
////	return postsService.postModel.FindOneByEmail(email)
////}
//
//func (postsService *postsService) AddMessageByPostId(account *models.Account, createPostDto *dtos.CreatePostDto) {
//	message := &models.Message{Account: account, Post: GetOnePost(id), Text: addMessageDto.Text, PublishedAt: time.Now()}
//	postsService.messageModel.Add(message)
//}
//
//func (postsService *postsService) RemoveMessageByPostId(id int) error {
//	if !postsService.messageModel.Exists(id) {
//		return errors.New("post not found")
//	}
//	postsService.messageModel.Delete(id)
//	return nil
//}
