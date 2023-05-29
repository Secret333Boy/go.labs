package posts

import (
	"errors"
	"go.labs/server/app/models"
	"reflect"
	"testing"
)

func TestGetPost(t *testing.T) {
	notFoundError := errors.New("post was not found error")
	post := models.Post{
		ID:          1,
		AccountID:   1,
		Title:       "test",
		Description: "test",
	}

	testCases := []struct {
		id             int
		description    string
		expectedResult *models.Post
		expectedError  error
	}{
		{
			id:             1,
			description:    "Post was found",
			expectedResult: &post,
		},
		{
			id:            2,
			description:   "Post was not found",
			expectedError: notFoundError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			postRepositoryMock := &PostRepositoryMock{
				Post:     tc.expectedResult,
				GetError: tc.expectedError,
			}

			p := NewPostsService(postRepositoryMock, &MessagesRepositoryMock{})
			result, err := p.GetOnePost(tc.id)

			if !reflect.DeepEqual(tc.expectedResult, result) {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedResult, result)
			}

			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}

func TestCreatePost(t *testing.T) {
	postCreateError := errors.New("failed to create post")
	testCases := []struct {
		description     string
		postCreateError error
		expectedError   error
	}{
		{
			description:     "get error from post repository",
			postCreateError: postCreateError,
			expectedError:   postCreateError,
		},
		{
			description: "success post create",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			postRepositoryMock := &PostRepositoryMock{
				CreateError: tc.postCreateError,
			}

			p := NewPostsService(postRepositoryMock, &MessagesRepositoryMock{})
			err := p.AddPost(&models.Account{}, "", "")
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	postDeleteError := errors.New("failed to delete post")
	post := models.Post{
		ID:          2,
		AccountID:   1,
		Title:       "test 2",
		Description: "test",
	}
	testCases := []struct {
		id              int
		description     string
		postDeleteError error
		expectedError   error
	}{
		{
			id:              1,
			description:     "get error from post repository",
			postDeleteError: postDeleteError,
			expectedError:   postDeleteError,
		},
		{
			id:          2,
			description: "success post delete",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			postRepositoryMock := &PostRepositoryMock{
				Post:        &post,
				DeleteError: tc.postDeleteError,
			}

			p := NewPostsService(postRepositoryMock, &MessagesRepositoryMock{})

			err := p.RemovePost(&models.Account{}, post.ID)
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}
