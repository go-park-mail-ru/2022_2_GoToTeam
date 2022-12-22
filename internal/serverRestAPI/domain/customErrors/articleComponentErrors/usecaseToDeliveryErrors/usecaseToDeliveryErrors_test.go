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

func TestArticleDoesntExistError(t *testing.T) {
	errMessage := "err123"
	re := ArticleDoesntExistError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestEmailForSessionDoesntExistError(t *testing.T) {
	errMessage := "err123"
	re := EmailForSessionDoesntExistError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}

func TestEmailIsNotAuthorError(t *testing.T) {
	errMessage := "err123"
	re := EmailIsNotAuthorError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}
