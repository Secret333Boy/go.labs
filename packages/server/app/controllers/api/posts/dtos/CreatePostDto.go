package dtos

import "errors"

type CreatePostDto struct {
	Title       string
	Description string
}

const NoTitleErrorMessage = "title is mandatory"

func (postsDto *CreatePostDto) Validate() error {
	if postsDto.Title == "" {
		return errors.New(NoTitleErrorMessage)
	}
	return nil
}
