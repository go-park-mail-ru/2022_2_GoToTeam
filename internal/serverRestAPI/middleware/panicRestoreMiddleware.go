package middleware

import (
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"fmt"
	"github.com/labstack/echo/v4"
)

func PanicRestoreMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				context.Error(errorsUtils.WrapError("error while panic restoring", fmt.Errorf("%s", err)))
			}
		}()

		return next(context)
	}
}
