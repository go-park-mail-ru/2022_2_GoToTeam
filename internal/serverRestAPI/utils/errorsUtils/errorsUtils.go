package errorsUtils

import "fmt"

func WrapError(wrappingErrorMessage string, err error) error {
	return fmt.Errorf("%s: %w", wrappingErrorMessage, err)
}
