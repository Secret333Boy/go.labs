package posts

import "go.labs/server/app/models"

type MessagesRepositoryMock struct {
	Message     *models.Message
	GetError    error
	CreateError error
	DeleteError error
}

func (m MessagesRepositoryMock) FindOneMessageByPostId(int, int) (*models.Message, error) {
	return m.Message, m.CreateError
}

func (m MessagesRepositoryMock) CreateMessage(*models.Account, int, string) error {
	return m.CreateError
}

func (m MessagesRepositoryMock) RemoveMessage(*models.Message, *models.Account, int, int) error {
	return m.DeleteError
}

func (m MessagesRepositoryMock) FindAllMessagesByPostId(int, int, int) ([]models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessagesRepositoryMock) UpdateMessage(*models.Message, *models.Account, int, int, string) error {
	//TODO implement me
	panic("implement me")
}
