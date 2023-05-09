package message

import (
	"errors"
	"fmt"
	"time"

	"go.labs/server/app/models"
	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(DB *gorm.DB) *MessageRepository {
	return &MessageRepository{DB: DB}
}

func (mr *MessageRepository) FindAllMessagesByPostId(postId int, limit int, offset int) ([]models.Message, error) {
	var messages []models.Message

	if result := mr.DB.Offset(offset).Limit(limit).Find(&messages); result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}

func (mr *MessageRepository) FindOneMessageByPostId(postId int, messageId int) (*models.Message, error) {
	message := &models.Message{}

	result := mr.DB.First(&message, messageId)
	if result.Error != nil {
		return nil, result.Error
	}
	//TODO: return 404
	return message, nil
}

func (mr *MessageRepository) CreateMessage(account *models.Account, postId int, text string) error {
	message := &models.Message{AccountID: account.ID, PostID: uint(postId), Text: text, PublishedAt: time.Now()}
	if result := mr.DB.Create(&message); result.Error != nil {
		return result.Error
	}
	return nil
}

func (mr *MessageRepository) UpdateMessage(message *models.Message, account *models.Account, postId int, messageId int, text string) error {
	if result := mr.DB.Find(&message, "ID = ? AND Account_ID = ? AND Post_ID = ?", messageId, account.ID, postId); result.Error != nil {
		fmt.Println(result.Error)
		return errors.New("access denied")
	}

	message.Text = text
	mr.DB.Save(&message)

	return nil
}

func (mr *MessageRepository) RemoveMessage(message *models.Message, account *models.Account, postId int, messageId int) error {
	if result := mr.DB.Find(&message, "ID = ? AND Account_ID = ? AND Post_ID = ?", messageId, account.ID, postId); result.Error != nil {
		fmt.Println(result.Error)
		return errors.New("access denied")
	}

	if result := mr.DB.Delete(message); result.Error != nil {
		fmt.Println(result.Error)
	}
	return nil
}
