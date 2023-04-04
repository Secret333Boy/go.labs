package models

import (
	"time"
)

type Message struct {
	Id          int
	Account     *Account
	Post        *Post
	Text        string
	PublishedAt time.Time
	model       *MessageModel
}

type MessageModel struct {
	messages []Message
	lastId   int
}

func NewMessageModel() *MessageModel {
	var model = new(MessageModel)
	model.messages = make([]Message, 0)
	return model
}

func (model *MessageModel) FindAll(postId int, limit int, offset int) []Message {

	var messagesByPostId = make([]Message, 0, len(model.messages))

	for _, message := range model.messages {
		if message.Post.Id == postId {
			messagesByPostId = append(messagesByPostId, message)
		}
	}

	if len(messagesByPostId) > 0 {
		if offset > len(messagesByPostId) {
			offset = len(messagesByPostId)
		}

		end := offset + limit
		if end > len(messagesByPostId) {
			end = len(messagesByPostId)
		}
		return messagesByPostId[offset:end]
	} else {
		return messagesByPostId
	}
}

func (model *MessageModel) FindOne(postId, id int) *Message {
	for _, message := range model.messages {
		if message.Id == id && message.Post.Id == postId {
			return &message
		}
	}
	return nil
}

func (model *MessageModel) Add(message *Message) {
	message.model = model
	message.Id = model.lastId + 1
	model.messages = append(model.messages, *message)
	model.lastId++
}

func (model *MessageModel) Delete(id int) {
	for i, message := range model.messages {
		if message.Id == id {
			message.model = nil
			model.messages = append(model.messages[:i], model.messages[i+1:]...)
			return
		}
	}
}

func (model *MessageModel) Exists(id int) bool {
	for _, message := range model.messages {
		if message.Id == id {
			return true
		}
	}
	return false
}
