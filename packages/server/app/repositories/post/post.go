package post

import (
	"errors"
	"fmt"
	"time"

	"go.labs/server/app/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(DB *gorm.DB) *PostRepository {
	return &PostRepository{DB: DB}
}

func (pr *PostRepository) FindAllPosts(limit int, offset int) []models.Post {
	var posts []models.Post
	if result := pr.DB.Offset(offset).Limit(limit).Find(&posts); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}
	return posts
}

func (pr *PostRepository) FindOnePost(id int) (*models.Post, error) {
	post := &models.Post{}
	result := pr.DB.First(post, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	//TODO: return 404
	return post, nil
}

func (pr *PostRepository) CreatePost(account *models.Account, title string, description string) error {
	post := &models.Post{AccountID: account.ID, Title: title, Description: description, PublishedAt: time.Now()}

	if result := pr.DB.Create(&post); result.Error != nil {
		fmt.Println(result.Error)
		return errors.New("error while creating post")
	}

	return nil
}

func (pr *PostRepository) UpdatePost(post *models.Post, account *models.Account, title string, description string) error {
	if post != nil {
		if result := pr.DB.Find(&post, "ID = ? AND Account_ID = ?", post.ID, account.ID); result.Error != nil {
			fmt.Println(result.Error)
			return errors.New("access denied")
		}
		post.Title = title
		post.Description = description
		pr.DB.Save(&post)
	} else {
		return errors.New("post not found")
	}
	return nil
}

func (pr *PostRepository) RemovePost(post *models.Post, account *models.Account) error {
	if post != nil {
		if result := pr.DB.Find(&post, "ID = ? AND Account_ID = ?", post.ID, account.ID); result.Error != nil {
			fmt.Println(result.Error)
			return errors.New("access denied")
		}
		if result := pr.DB.Delete(post); result.Error != nil {
			fmt.Println(result.Error)
		}
	} else {
		return errors.New("post not found")
	}
	return nil
}
