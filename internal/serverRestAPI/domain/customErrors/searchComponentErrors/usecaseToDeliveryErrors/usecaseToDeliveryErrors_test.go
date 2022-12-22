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

func TestTagDoesntExistError(t *testing.T) {
	errMessage := "err123"
	re := TagDoesntExistError{Err: errors.New(errMessage)}
	assert.Equal(t, errMessage, re.Error())
}
