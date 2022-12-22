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

func TestEmailDoesntExistError(t *testing.T) {
	errMessage := "err123"
	re := EmailDoesntExistError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestEmailExistsError(t *testing.T) {
	errMessage := "err123"
	re := EmailExistsError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestLoginExistsError(t *testing.T) {
	errMessage := "err123"
	re := LoginExistsError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestEmailIsNotValidError(t *testing.T) {
	errMessage := "err123"
	re := EmailIsNotValidError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestLoginIsNotValidError(t *testing.T) {
	errMessage := "err123"
	re := LoginIsNotValidError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestPasswordIsNotValidError(t *testing.T) {
	errMessage := "err123"
	re := PasswordIsNotValidError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}
