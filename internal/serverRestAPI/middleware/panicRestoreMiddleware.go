package middleware

import (
	"2022_2_GoTo_team/pkg/utils/logger"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PanicRestoreMiddleware(logger *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			defer func(ctx echo.Context) {
				if err := recover(); err != nil {
					logger.LogrusLoggerWithContext(ctx.Request().Context()).Error("Enter to the panic restore middleware defer function. Error: ", fmt.Errorf("%s", err), ". Request: ", ctx.Request())
				}

				err := ctx.NoContent(http.StatusInternalServerError)
				if err != nil {
					logger.LogrusLoggerWithContext(ctx.Request().Context()).Error("Error in the panic restore middleware defer function. Error: ", err)
				}
			}(ctx)

			return next(ctx)
		}
	}
}
