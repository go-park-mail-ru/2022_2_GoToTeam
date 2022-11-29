package errorsUtils

import (
	"fmt"
	"google.golang.org/grpc/status"
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
		return 500
	}
	if statusCode == 2 {
		return 500
	}

	return statusCode
}
