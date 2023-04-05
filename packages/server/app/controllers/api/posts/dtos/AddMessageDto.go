package dtos

import "errors"

type AddMessageDto struct {
	Text string
}

const NoTextErrorMessage = "Add message text"

func (addMessageDto *AddMessageDto) Validate() error {
	if addMessageDto.Text == "" {
		return errors.New(NoTextErrorMessage)
	}
	return nil
}
