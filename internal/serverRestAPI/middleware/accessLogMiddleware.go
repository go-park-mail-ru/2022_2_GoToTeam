package middleware

import (
	"2022_2_GoTo_team/pkg/domain"
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
				requestMethod := ctx.Request().Method
				urlPath := ctx.Request().URL.Path
				responseCode := ctx.Response().Status
				elapsedTime := time.Since(requestProcessStartTime).Seconds()

				RecordHits(requestMethod, urlPath, responseCode)
				RecordLatency(requestMethod, urlPath, elapsedTime)

				logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("Request process finished. Elapsed time: ", elapsedTime, " seconds.")
			}()
			//c.Request().Header.Set(echo.HeaderXRequestID, uuid.New().String())
			ctx.SetRequest(ctx.Request().Clone(context.WithValue(ctx.Request().Context(), domain.REQUEST_ID_KEY_FOR_CONTEXT, uuid.New().String())))

			r := ctx.Request()
			logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("Request method: ", r.Method, ", remote address: ", r.RemoteAddr, ", request URL: ", r.URL.Path, ", request process start time: ", requestProcessStartTime)

			return next(ctx)
		}
	}
}
