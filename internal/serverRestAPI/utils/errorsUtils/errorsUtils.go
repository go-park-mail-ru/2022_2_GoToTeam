package errorsUtils

import "fmt"

func WrapError(wrappingErrorMessage string, wrappedError error) error {
	return fmt.Errorf("%s: %w", wrappingErrorMessage, wrappedError)
}
