package posts

import (
	"errors"
	"go.labs/server/app/models"
	"reflect"
	"testing"
	"time"
)

func TestGetMessage(t *testing.T) {
	createError := errors.New("message was not found")
	post := models.Post{
		ID:          1,
		AccountID:   1,
		Title:       "test",
		Description: "test",
	}
	message := models.Message{
		ID:          1,
		AccountID:   1,
		PostID:      1,
		Text:        "Test message",
		PublishedAt: time.Time{},
	}

	testCases := []struct {
		id             int
		description    string
		expectedResult *models.Message
		expectedError  error
	}{
		{
			id:             1,
			description:    "Message was found",
			expectedResult: &message,
		},
		{
			id:            2,
			description:   "Message was not found",
			expectedError: createError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			postRepositoryMock := &PostRepositoryMock{
				Post: &post,
			}

			messageRepositoryMock := &MessagesRepositoryMock{
				Message:     tc.expectedResult,
				CreateError: tc.expectedError,
			}

			p := NewPostsService(postRepositoryMock, messageRepositoryMock)
			result, err := p.GetOneMessageByPostId(post.ID, tc.id)

			if !reflect.DeepEqual(tc.expectedResult, result) {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedResult, result)
			}

			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}

func TestCreateMessage(t *testing.T) {
	messageCreateError := errors.New("failed to create message")
	post := models.Post{
		ID:          1,
		AccountID:   1,
		Title:       "test",
		Description: "test",
	}
	testCases := []struct {
		description        string
		messageCreateError error
		expectedError      error
	}{
		{
			description:        "get error from message repository",
			messageCreateError: messageCreateError,
			expectedError:      messageCreateError,
		},
		{
			description: "success message create",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			postRepositoryMock := &PostRepositoryMock{
				Post: &post,
			}
			messageRepositoryMock := &MessagesRepositoryMock{
				CreateError: tc.messageCreateError,
			}

			p := NewPostsService(postRepositoryMock, messageRepositoryMock)
			err := p.AddMessageByPostId(&models.Account{}, post.ID, "test")
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}

func TestDeleteMessage(t *testing.T) {
	messageDeleteError := errors.New("failed to delete message")
	post := models.Post{
		ID:          2,
		AccountID:   1,
		Title:       "test 2",
		Description: "test",
	}
	message := models.Message{
		ID:          1,
		AccountID:   1,
		PostID:      1,
		Text:        "Test message",
		PublishedAt: time.Time{},
	}
	testCases := []struct {
		id                 int
		description        string
		messageDeleteError error
		expectedError      error
	}{
		{
			id:                 1,
			description:        "get error from message repository",
			messageDeleteError: messageDeleteError,
			expectedError:      messageDeleteError,
		},
		{
			id:          2,
			description: "success message delete",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			postRepositoryMock := &PostRepositoryMock{
				Post: &post,
			}
			messageRepositoryMock := &MessagesRepositoryMock{
				Message:     &message,
				DeleteError: tc.messageDeleteError,
			}

			p := NewPostsService(postRepositoryMock, messageRepositoryMock)

			err := p.RemoveMessageByPostId(&models.Account{}, post.ID, tc.id)
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}
