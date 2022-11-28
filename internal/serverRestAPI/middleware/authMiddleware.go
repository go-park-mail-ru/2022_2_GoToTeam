package middleware

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/logger"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	needAuthUrls = map[string]struct{}{
		"/api/v1/session/info":   {},
		"/api/v1/session/remove": {},
		"/api/v1/article/create": {},
		"/api/v1/article/remove": {},

		"/api/v1/commentary/create": {},
	}
	noNeedSessionUrls = map[string]struct{}{
		"/": struct{}{},
	}
)

func isAuthorized(ctx echo.Context, sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface, logger *logger.Logger) bool {
	logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("Enter to the isAuthorized function.")

	authorized := false
	if cookie, err := ctx.Cookie(domain.SESSION_COOKIE_HEADER_NAME); err == nil && cookie != nil {
		if authorized, err = sessionUsecase.SessionExists(ctx.Request().Context(), &models.Session{SessionId: cookie.Value}); err != nil {
			return false
		}
	}

	return authorized
}

func getCookieValue(ctx echo.Context) *http.Cookie {
	cookie, err := ctx.Cookie(domain.SESSION_COOKIE_HEADER_NAME)
	if err != nil {
		return nil
	}

	return cookie
}

func AuthMiddleware(sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface, logger *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("Auth middleware. URL.Path: ", ctx.Request().URL.Path)

			if _, ok := needAuthUrls[ctx.Request().URL.Path]; !ok {
				logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("Dont need auth for the URL.Path: ", ctx.Request().URL.Path)
				return next(ctx)
			}

			if !isAuthorized(ctx, sessionUsecase, logger) {
				logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("Unauthorized!")
				return ctx.NoContent(http.StatusUnauthorized)
			}

			cookie := getCookieValue(ctx)
			if cookie != nil {
				email, err := sessionUsecase.GetUserEmailBySession(context.Background(), &models.Session{SessionId: cookie.Value})
				if err != nil {
					logger.LogrusLoggerWithContext(ctx.Request().Context()).Error(err)
				}

				ctx.SetRequest(ctx.Request().Clone(context.WithValue(ctx.Request().Context(), domain.USER_EMAIL_KEY_FOR_CONTEXT, email)))
			}

			logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("Authorized!")

			return next(ctx)
		}
	}
}
