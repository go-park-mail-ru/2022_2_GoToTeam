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

			logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("A")
			csrf := r.Header.Get("X-XSRF-Token")
			logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("B")
			k, _ := r.Cookie("_csrf")
			logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("C")

			logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("X-XSRF-Token header = ", csrf, " _csrf cookie = ", k)
			//assert.Equal(k, csrf)
			if csrf == "" || k == nil {
				logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("CSRF Security failed. csrf == '' or k == nil")
			} else if csrf != k.Value {
				logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("CSRF Security failed.")
			} else {
				logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("CSRF Security ok")
			}

			return next(ctx)
		}
	}
}
