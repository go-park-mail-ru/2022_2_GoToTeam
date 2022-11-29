package middleware

import (
	"2022_2_GoTo_team/pkg/domain/constants"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"
)

func AccessLogMiddleware(logger *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			requestProcessStartTime := time.Now()

			defer func() {
				logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("Request process finished. Spent time: ", time.Since(requestProcessStartTime))
			}()
			//c.Request().Header.Set(echo.HeaderXRequestID, uuid.New().String())
			ctx.SetRequest(ctx.Request().Clone(context.WithValue(ctx.Request().Context(), constants.REQUEST_ID_KEY_FOR_CONTEXT, uuid.New().String())))

			r := ctx.Request()
			logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("Request method: ", r.Method, ", remote address: ", r.RemoteAddr, ", request URL: ", r.URL.Path, ", request process start time: ", requestProcessStartTime)

			return next(ctx)
		}
	}
}
