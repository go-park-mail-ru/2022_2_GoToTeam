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

func TestGetFeedError(t *testing.T) {
	errMessage := "err123"
	re := GetFeedError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestLoginIsNotValidError(t *testing.T) {
	errMessage := "err123"
	re := LoginIsNotValidError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestLoginDoesntExistError(t *testing.T) {
	errMessage := "err123"
	re := LoginDoesntExistError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestCategoryDoesntExistError(t *testing.T) {
	errMessage := "err123"
	re := CategoryDoesntExistError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}
