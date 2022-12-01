package errorsUtils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
	"testing"
)

func TestWrapError(t *testing.T) {
	err := errors.New("custom error")
	wrappingStr := "wm"
	newErr := WrapError(wrappingStr, err)
	assert.Equal(t, err, errors.Unwrap(newErr))
	if !strings.HasPrefix(newErr.Error(), wrappingStr) {
		t.Error("incorrect wrapping")
	}
}

func TestExtractCodeFromGrpcErrorStatus(t *testing.T) {
	st := &status.Status{}
	res := ExtractCodeFromGrpcErrorStatus(st)
	assert.Equal(t, http.StatusInternalServerError, res)
}
