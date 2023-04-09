package services

import (
	"errors"
	"go.labs/server/app/controllers/api/posts"
	"time"

	"go.labs/server/app/models"
)

type postsService struct {
	postModel    *models.PostModel
	messageModel *models.MessageModel
}

func (p *postsService) GetAllPosts(limit int, offset int) []models.Post {
	return p.postModel.FindAll(limit, offset)
}

func (p *postsService) GetOnePost(id int) *models.Post {
	return p.postModel.FindOne(id)
}

func (p *postsService) AddPost(account *models.Account, createPostDto *posts.CreatePostDto) {
	post := &models.Post{Account: account, Title: createPostDto.Title, Description: createPostDto.Description, PublishedAt: time.Now()}
	p.postModel.Add(post)
}

func (p *postsService) GetAccountByPostId(id int) *models.Account {
	return p.postModel.FindAccountByPostId(id)
}

func (p *postsService) UpdatePost(account *models.Account, id int, createPostDto *posts.CreatePostDto) error {
	if !p.postModel.Exists(id) {
		return errors.New("post not found")
	}
	if !(p.GetAccountByPostId(id).Id == account.Id) {
		return errors.New("access denied")
	}
	post := &models.Post{Title: createPostDto.Title, Description: createPostDto.Description}
	p.postModel.Update(id, post)
	return nil
}

func (p *postsService) RemovePost(account *models.Account, id int) error {
	if !p.postModel.Exists(id) {
		return errors.New("post not found")
	}
	if !(p.GetAccountByPostId(id).Id == account.Id) {
		return errors.New("access denied")
	}
	p.postModel.Delete(id)
	return nil
}

func (p *postsService) GetAllMessagesByPostId(postId int, limit int, offset int) ([]models.Message, error) {
	if !p.postModel.Exists(postId) {
		return nil, errors.New("post not found")
	}
	return p.messageModel.FindAll(postId, limit, offset), nil
}

func (p *postsService) GetOneMessageByPostId(postId int, messageId int) (*models.Message, error) {
	if !p.postModel.Exists(postId) {
		return nil, errors.New("post not found")
	}
	return p.messageModel.FindOne(postId, messageId), nil
}

func (p *postsService) AddMessageByPostId(account *models.Account, postId int, addMessageDto *posts.AddMessageDto) error {
	if !p.postModel.Exists(postId) {
		return errors.New("post not found")
	}
	message := &models.Message{Account: account, Post: p.GetOnePost(postId), Text: addMessageDto.Text, PublishedAt: time.Now()}
	p.messageModel.Add(message)
	return nil
}

func (p *postsService) UpdateMessageByPostId(account *models.Account, postId int, messageId int, addMessageDto *posts.AddMessageDto) error {
	if !p.postModel.Exists(postId) {
		return errors.New("post not found")
	}
	if !(p.GetAccountByPostId(postId).Id == account.Id) {
		return errors.New("access denied")
	}
	message := &models.Message{Text: addMessageDto.Text}
	p.messageModel.Update(postId, messageId, message)
	return nil
}

func (p *postsService) RemoveMessageByPostId(account *models.Account, postId int, messageId int) error {
	if !p.postModel.Exists(postId) {
		return errors.New("post not found")
	}
	if !(p.GetAccountByPostId(postId).Id == account.Id) {
		return errors.New("access denied")
	}
	if !p.messageModel.Exists(messageId) {
		return errors.New("message not found")
	}
	p.messageModel.Delete(messageId)
	return nil
}
