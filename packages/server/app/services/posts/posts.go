package posts

import (
	"errors"
	"time"

	"go.labs/server/app/models"
)

type PostsService struct {
	postModel    *models.PostModel
	messageModel *models.MessageModel
}

func (p *PostsService) GetAllPosts(limit int, offset int) []models.Post {
	return p.postModel.FindAll(limit, offset)
}

func (p *PostsService) GetOnePost(id int) *models.Post {
	return p.postModel.FindOne(id)
}

func (p *PostsService) AddPost(account *models.Account, title string, description string) {
	post := &models.Post{Account: account, Title: title, Description: description, PublishedAt: time.Now()}
	p.postModel.Add(post)
}

func (p *PostsService) GetAccountByPostId(id int) *models.Account {
	return p.postModel.FindAccountByPostId(id)
}

func (p *PostsService) UpdatePost(account *models.Account, id int, title string, description string) error {
	if !p.postModel.Exists(id) {
		return errors.New("post not found")
	}
	if !(p.GetAccountByPostId(id).Id == account.Id) {
		return errors.New("access denied")
	}
	post := &models.Post{Title: title, Description: description}
	p.postModel.Update(id, post)
	return nil
}

func (p *PostsService) RemovePost(account *models.Account, id int) error {
	if !p.postModel.Exists(id) {
		return errors.New("post not found")
	}
	if !(p.GetAccountByPostId(id).Id == account.Id) {
		return errors.New("access denied")
	}
	p.postModel.Delete(id)
	return nil
}

func (p *PostsService) GetAllMessagesByPostId(postId int, limit int, offset int) ([]models.Message, error) {
	if !p.postModel.Exists(postId) {
		return nil, errors.New("post not found")
	}
	return p.messageModel.FindAll(postId, limit, offset), nil
}

func (p *PostsService) GetOneMessageByPostId(postId int, messageId int) (*models.Message, error) {
	if !p.postModel.Exists(postId) {
		return nil, errors.New("post not found")
	}
	return p.messageModel.FindOne(postId, messageId), nil
}

func (p *PostsService) AddMessageByPostId(account *models.Account, postId int, text string) error {
	if !p.postModel.Exists(postId) {
		return errors.New("post not found")
	}
	message := &models.Message{Account: account, Post: p.GetOnePost(postId), Text: text, PublishedAt: time.Now()}
	p.messageModel.Add(message)
	return nil
}

func (p *PostsService) UpdateMessageByPostId(account *models.Account, postId int, messageId int, text string) error {
	if !p.postModel.Exists(postId) {
		return errors.New("post not found")
	}
	if !(p.GetAccountByPostId(postId).Id == account.Id) {
		return errors.New("access denied")
	}
	message := &models.Message{Text: text}
	p.messageModel.Update(postId, messageId, message)
	return nil
}

func (p *PostsService) RemoveMessageByPostId(account *models.Account, postId int, messageId int) error {
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

var postsServiceInstance = &PostsService{models.NewPostModel(), models.NewMessageModel()}

func GetPostsServiceInstance() *PostsService {
	return postsServiceInstance
}
