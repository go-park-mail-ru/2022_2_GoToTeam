package usecaseToDeliveryErrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryError(t *testing.T) {
	errMessage := "err123"
	re := RepositoryError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestIncorrectEmailOrPasswordError(t *testing.T) {
	errMessage := "err123"
	re := IncorrectEmailOrPasswordError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestEmailForSessionDoesntExistError(t *testing.T) {
	errMessage := "err123"
	re := EmailForSessionDoesntExistError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestUserForSessionDoesntExistError(t *testing.T) {
	errMessage := "err123"
	re := UserForSessionDoesntExistError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestEmailIsNotValidError(t *testing.T) {
	errMessage := "err123"
	re := EmailIsNotValidError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestPasswordIsNotValidError(t *testing.T) {
	errMessage := "err123"
	re := PasswordIsNotValidError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}
