package auth

import "errors"

type RegisterDto struct {
	Email    string
	Password string
}

const NoPasswordErrorMessage = "password must be defined"
const NoEmailErrorMessage = "email must be defined"

func (registerDto *RegisterDto) Validate() error {
	if registerDto.Email == "" {
		return errors.New(NoEmailErrorMessage)
	}

	if registerDto.Password == "" {
		return errors.New(NoPasswordErrorMessage)
	}

	return nil
}
