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

func TestOpenFileError(t *testing.T) {
	errMessage := "err123"
	re := OpenFileError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestFileSizeBigError(t *testing.T) {
	errMessage := "err123"
	re := FileSizeBigError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestNotImageError(t *testing.T) {
	errMessage := "err123"
	re := NotImageError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestEmailForSessionDoesntExistError(t *testing.T) {
	errMessage := "err123"
	re := EmailForSessionDoesntExistError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}
