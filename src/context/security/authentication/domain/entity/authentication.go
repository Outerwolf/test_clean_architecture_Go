package entity

import (
	"errors"
	"fmt"
	"github.com/auth/src/context/shared/domain/value_object"
)

var ErrInvalidAuthenticationID = errors.New("invalid authentication id")

type AuthenticationID struct {
	value string
}

func NewAuthenticationID(value string) (AuthenticationID, error) {
	v, err := value_object.ValidateUuid(value)
	if err != nil {
		return AuthenticationID{}, fmt.Errorf("%w: %s", ErrInvalidAuthenticationID, v)
	}

	return AuthenticationID{value: v}, nil

}

func (id AuthenticationID) String() string {
	return id.value
}

var ErrInvalidUsername = errors.New("invalid username")

type Username struct {
	value string
}

func NewUsername(value string) (Username, error) {
	if value == "" {
		return Username{}, fmt.Errorf("%w: %s", ErrInvalidUsername, value)
	}

	return Username{value: value}, nil
}

func (usr Username) String() string {
	return usr.value
}

var ErrInvalidPassword = errors.New("invalid password")
var ErrInvalidPasswordLength = errors.New("invalid password length")
var passMinLength = 8
var passMaxLength = 32

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	if value == "" {
		return Password{}, fmt.Errorf("%w: %s", ErrInvalidPassword, value)
	}

	if len(value) < passMinLength {
		return Password{}, fmt.Errorf("%w: %s", ErrInvalidPasswordLength, passMinLength)
	}

	if len(value) > passMaxLength {
		return Password{}, fmt.Errorf("%w: %s", ErrInvalidPasswordLength, passMaxLength)
	}
	return Password{value: value}, nil
}

func (pass Password) String() string {
	return pass.value
}

type Authentication struct {
	Id       AuthenticationID
	Username Username
	Password Password
}

func NewAuthentication(id string, email string, password string) (Authentication, error) {
	idVo, err := NewAuthenticationID(id)
	if err != nil {
		return Authentication{}, err
	}

	emailVo, err := NewUsername(email)

	if err != nil {

		return Authentication{}, err
	}

	passVo, err := NewPassword(password)

	if err != nil {

		return Authentication{}, err
	}

	return Authentication{
		Id:       idVo,
		Username: emailVo,
		Password: passVo,
	}, nil
}

func (auth Authentication) ID() AuthenticationID {
	return auth.Id
}

func (auth Authentication) GetEmail() Username {
	return auth.Username
}

func (auth Authentication) GetPassword() Password {
	return auth.Password
}
