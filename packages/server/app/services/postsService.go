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

func (postsService *postsService) GetAccountByPostId(id int) *models.Account {
	return postsService.postModel.FindAccountByPostId(id)
}

func (postsService *postsService) UpdatePost(account *models.Account, id int, createPostDto *dtos.CreatePostDto) error {
	if !postsService.postModel.Exists(id) {
		return errors.New("post not found")
	}
	if !(postsService.GetAccountByPostId(id).Id == account.Id) {
		return errors.New("access denied")
	}
	post := &models.Post{Title: createPostDto.Title, Description: createPostDto.Description}
	postsService.postModel.Update(id, post)
	return nil
}

func (postsService *postsService) RemovePost(account *models.Account, id int) error {
	if !postsService.postModel.Exists(id) {
		return errors.New("post not found")
	}
	if !(postsService.GetAccountByPostId(id).Id == account.Id) {
		return errors.New("access denied")
	}
	postsService.postModel.Delete(id)
	return nil
}

func (postsService *postsService) GetAllMessagesByPostId(postId int, limit int, offset int) []models.Message {
	return postsService.messageModel.FindAll(postId, limit, offset)
}

func (postsService *postsService) GetOneMessageByPostId(postId int, messageId int) *models.Message {
	return postsService.messageModel.FindOne(postId, messageId)
}

func (postsService *postsService) AddMessageByPostId(account *models.Account, postId int, addMessageDto *dtos.AddMessageDto) {
	message := &models.Message{Account: account, Post: postsService.GetOnePost(postId), Text: addMessageDto.Text, PublishedAt: time.Now()}
	postsService.messageModel.Add(message)
}

func (postsService *postsService) UpdateMessageByPostId(account *models.Account, postId int, messageId int, addMessageDto *dtos.AddMessageDto) error {
	if !postsService.postModel.Exists(postId) {
		return errors.New("post not found")
	}
	if !(postsService.GetAccountByPostId(postId).Id == account.Id) {
		return errors.New("access denied")
	}
	message := &models.Message{Text: addMessageDto.Text}
	postsService.messageModel.Update(postId, messageId, message)
	return nil
}

func (postsService *postsService) RemoveMessageByPostId(account *models.Account, postId int, messageId int) error {
	if !postsService.postModel.Exists(postId) {
		return errors.New("post not found")
	}
	if !(postsService.GetAccountByPostId(postId).Id == account.Id) {
		return errors.New("access denied")
	}
	if !postsService.messageModel.Exists(messageId) {
		return errors.New("message not found")
	}
	postsService.messageModel.Delete(messageId)
	return nil
}
