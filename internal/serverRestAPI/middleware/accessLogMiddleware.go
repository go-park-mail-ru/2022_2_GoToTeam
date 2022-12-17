package middleware

import (
	"2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func AccessLogMiddleware(logger *logger.Logger, enableEchoCsrfToken bool) echo.MiddlewareFunc {
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

			if enableEchoCsrfToken {
				csrf_value := ctx.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
				logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("_csrf_value from config context = ", csrf_value)

				if r.Method == "POST" || r.Method == "PUT" {
					logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("CSRF token validation begin. Method: ", r.Method)
					csrfHeader := r.Header.Get("X-XSRF-Token")
					csrfCookie, _ := r.Cookie("_csrf")
					logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("X-XSRF-Token csrfHeader = ", csrfHeader, " csrfCookie = ", csrfCookie)

					//assert.Equal(k, csrf)
					if csrfHeader == "" || csrfCookie == nil {
						logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("CSRF Security failed. csrfHeader is empty or csrfCookie == nil.")
						//return ctx.NoContent(http.StatusForbidden)
					} else if csrfHeader != csrfCookie.Value {
						logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("CSRF Security failed. csrfHeader != csrfCookie.Value")
						//return ctx.NoContent(http.StatusForbidden)
					}

					logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("CSRF security successfully validated.")
				} else {
					logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("Dont need CSRF token validation. Method: ", r.Method)
				}
			}

			return next(ctx)
		}
	}
}
