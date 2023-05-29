package posts

import (
	"errors"

	"go.labs/server/app/models"
)

type postRepository interface {
	FindAllPosts(limit int, offset int) []models.Post
	FindOnePost(id int) (*models.Post, error)
	CreatePost(account *models.Account, title string, description string) error
	UpdatePost(post *models.Post, account *models.Account, title string, description string) error
	RemovePost(post *models.Post, account *models.Account) error
}
type messageRepository interface {
	FindAllMessagesByPostId(postId int, limit int, offset int) ([]models.Message, error)
	FindOneMessageByPostId(postId int, messageId int) (*models.Message, error)
	CreateMessage(account *models.Account, postId int, text string) error
	UpdateMessage(message *models.Message, account *models.Account, postId int, messageId int, text string) error
	RemoveMessage(message *models.Message, account *models.Account, postId int, messageId int) error
}

type postsService struct {
	postRepository    postRepository
	messageRepository messageRepository
}

func NewPostsService(postRepository postRepository, messageRepository messageRepository) Posts {
	return &postsService{postRepository: postRepository, messageRepository: messageRepository}
}

type Posts interface {
	GetAllPosts(limit int, offset int) []models.Post
	GetOnePost(id int) (*models.Post, error)
	AddPost(account *models.Account, title string, description string) error
	UpdatePost(account *models.Account, id int, title string, description string) error
	RemovePost(account *models.Account, id int) error

	GetAllMessagesByPostId(postId int, limit int, offset int) ([]models.Message, error)
	GetOneMessageByPostId(postId int, messageId int) (*models.Message, error)
	AddMessageByPostId(account *models.Account, postId int, text string) error
	UpdateMessageByPostId(account *models.Account, postId int, messageId int, text string) error
	RemoveMessageByPostId(account *models.Account, postId int, messageId int) error
}

func (p *postsService) GetAllPosts(limit int, offset int) []models.Post {
	return p.postRepository.FindAllPosts(limit, offset)
}

func (p *postsService) GetOnePost(id int) (*models.Post, error) {
	return p.postRepository.FindOnePost(id)
}

func (p *postsService) AddPost(account *models.Account, title string, description string) error {
	return p.postRepository.CreatePost(account, title, description)
}

func (p *postsService) UpdatePost(account *models.Account, id int, title string, description string) error {
	post, _ := p.GetOnePost(id)
	return p.postRepository.UpdatePost(post, account, title, description)
}

func (p *postsService) RemovePost(account *models.Account, id int) error {
	post, _ := p.GetOnePost(id)
	return p.postRepository.RemovePost(post, account)
}

func (p *postsService) GetAllMessagesByPostId(postId int, limit int, offset int) ([]models.Message, error) {

	post, _ := p.GetOnePost(postId)

	if post == nil {
		return nil, errors.New("This post is not exist")
	}

	return p.messageRepository.FindAllMessagesByPostId(postId, limit, offset)
}

func (p *postsService) GetOneMessageByPostId(postId int, messageId int) (*models.Message, error) {
	post, _ := p.GetOnePost(postId)

	if post == nil {
		return nil, errors.New("This post is not exist")
	}

	return p.messageRepository.FindOneMessageByPostId(postId, messageId)
}

func (p *postsService) AddMessageByPostId(account *models.Account, postId int, text string) error {
	post, _ := p.GetOnePost(postId)

	if post == nil {
		return errors.New("This post is not exist")

	}

	return p.messageRepository.CreateMessage(account, postId, text)
}

func (p *postsService) UpdateMessageByPostId(account *models.Account, postId int, messageId int, text string) error {
	post, _ := p.GetOnePost(postId)

	if post == nil {
		return errors.New("post not found")
	}

	message, err := p.GetOneMessageByPostId(postId, messageId)
	if err != nil {
		return errors.New("message not found")
	}

	return p.messageRepository.UpdateMessage(message, account, postId, messageId, text)
}

func (p *postsService) RemoveMessageByPostId(account *models.Account, postId int, messageId int) error {
	post, _ := p.GetOnePost(postId)

	if post == nil {
		return errors.New("post not found")
	}

	message, err := p.GetOneMessageByPostId(postId, messageId)
	if err != nil {
		return errors.New("message not found")
	}

	return p.messageRepository.RemoveMessage(message, account, postId, messageId)
}
