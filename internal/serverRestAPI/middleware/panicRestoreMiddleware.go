package middleware

import (
	"2022_2_GoTo_team/pkg/logger"
	"fmt"
	"github.com/labstack/echo/v4"
)

func PanicRestoreMiddleware(logger *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					logger.LogrusLoggerWithContext(ctx.Request().Context()).Error("Enter to the panic restore middleware defer function. Error: ", fmt.Errorf("%s", err), ". Request: ", ctx.Request())
				}
			}()

			return next(ctx)
		}
	}
}
