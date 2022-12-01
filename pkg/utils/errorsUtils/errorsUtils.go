package errorsUtils

import (
	"fmt"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
)

func WrapError(wrappingErrorMessage string, wrappedError error) error {
	return fmt.Errorf("%s: %w", wrappingErrorMessage, wrappedError)
}

func ExtractCodeFromGrpcErrorStatus(status *status.Status) int {
	statusCodeStr := status.Code().String()
	statusCodeStr = strings.TrimLeft(statusCodeStr, "Code(")
	statusCodeStr = strings.TrimRight(statusCodeStr, ")")

	statusCode, err := strconv.Atoi(statusCodeStr)
	if err != nil {
		return http.StatusInternalServerError
	}
	if statusCode == 2 {
		return http.StatusInternalServerError
	}

	return statusCode
}
