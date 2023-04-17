package posts

import "errors"

type AddMessageDto struct {
	Text string
}

const NoTextErrorMessage = "add message text"

func (addMessageDto *AddMessageDto) Validate() error {
	if addMessageDto.Text == "" {
		return errors.New(NoTextErrorMessage)
	}
	return nil
}
