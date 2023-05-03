package posts

import (
	"errors"
	"fmt"
	"go.labs/server/app/models"
	"gorm.io/gorm"
	"time"
)

type postsService struct {
	DB *gorm.DB
}

func NewPostsService(db *gorm.DB) PostsAndMessages {
	return &postsService{DB: db}
}

type PostsAndMessages interface {
	GetAllPosts(limit int, offset int) []models.Post
	GetOnePost(id int) *models.Post
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
	var posts []models.Post
	if result := p.DB.Offset(offset).Limit(limit).Find(&posts); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}
	return posts
}

func (p *postsService) GetOnePost(id int) *models.Post {
	post := &models.Post{}
	result := p.DB.First(post, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}
	//TODO: return 404
	return post
}

func (p *postsService) AddPost(account *models.Account, title string, description string) error {
	post := &models.Post{AccountID: account.ID, Title: title, Description: description, PublishedAt: time.Now()}

	if result := p.DB.Create(&post); result.Error != nil {
		fmt.Println(result.Error)
		return errors.New("error while creating post")
	}

	return nil
}

func (p *postsService) UpdatePost(account *models.Account, id int, title string, description string) error {
	post := p.GetOnePost(id)
	if post != nil {
		if result := p.DB.Find(&post, "ID = ? AND Account_ID = ?", id, account.ID); result.Error != nil {
			fmt.Println(result.Error)
			return errors.New("access denied")
		}
		post.Title = title
		post.Description = description
		p.DB.Save(&post)
	} else {
		return errors.New("post not found")
	}
	return nil
}

func (p *postsService) RemovePost(account *models.Account, id int) error {
	post := p.GetOnePost(id)
	if post != nil {
		if result := p.DB.Find(&post, "ID = ? AND Account_ID = ?", id, account.ID); result.Error != nil {
			fmt.Println(result.Error)
			return errors.New("access denied")
		}
		if result := p.DB.Delete(post); result.Error != nil {
			fmt.Println(result.Error)
		}
	} else {
		return errors.New("post not found")
	}
	return nil
}

func (p *postsService) GetAllMessagesByPostId(postId int, limit int, offset int) ([]models.Message, error) {

	post := p.GetOnePost(postId)
	var messages []models.Message

	if post != nil {
		if result := p.DB.Offset(offset).Limit(limit).Find(&messages); result.Error != nil {
			return nil, result.Error
		}
	} else {
		return nil, errors.New("This post is not exist")
	}
	return messages, nil
}

func (p *postsService) GetOneMessageByPostId(postId int, messageId int) (*models.Message, error) {

	post := p.GetOnePost(postId)
	message := &models.Message{}

	if post != nil {
		result := p.DB.First(&message, messageId)
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		return nil, errors.New("This post is not exist")
	}
	//TODO: return 404
	return message, nil
}

func (p *postsService) AddMessageByPostId(account *models.Account, postId int, text string) error {

	post := p.GetOnePost(postId)

	if post != nil {
		message := &models.Message{AccountID: account.ID, PostID: uint(postId), Text: text, PublishedAt: time.Now()}
		if result := p.DB.Create(&message); result.Error != nil {
			return result.Error
		}
	} else {
		return errors.New("This post is not exist")
	}
	return nil
}

func (p *postsService) UpdateMessageByPostId(account *models.Account, postId int, messageId int, text string) error {

	post := p.GetOnePost(postId)

	if post != nil {

		message, err := p.GetOneMessageByPostId(postId, messageId)
		if err == nil {
			if result := p.DB.Find(&message, "ID = ? AND Account_ID = ? AND Post_ID = ?", messageId, account.ID, postId); result.Error != nil {
				fmt.Println(result.Error)
				return errors.New("access denied")
			}
			message.Text = text
			p.DB.Save(&message)
		} else {
			return errors.New("message not found")
		}

	} else {
		return errors.New("post not found")
	}
	return nil

}

func (p *postsService) RemoveMessageByPostId(account *models.Account, postId int, messageId int) error {

	post := p.GetOnePost(postId)

	if post != nil {
		message, err := p.GetOneMessageByPostId(postId, messageId)
		if err == nil {
			if result := p.DB.Find(&message, "ID = ? AND Account_ID = ? AND Post_ID = ?", messageId, account.ID, postId); result.Error != nil {
				fmt.Println(result.Error)
				return errors.New("access denied")
			}
			if result := p.DB.Delete(message); result.Error != nil {
				fmt.Println(result.Error)
			}
		} else {
			return errors.New("message not found")
		}
	} else {
		return errors.New("post not found")
	}
	return nil
}
