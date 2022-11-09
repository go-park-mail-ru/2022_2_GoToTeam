package middleware

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	needAuthUrls = map[string]struct{}{
		"/api/v1/article/create": struct{}{},
	}
	noNeedSessionUrls = map[string]struct{}{
		"/": struct{}{},
	}
)

func isAuthorized(c echo.Context, sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface, logger *logger.Logger) bool {
	logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the isAuthorized function.")

	authorized := false
	if cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME); err == nil && cookie != nil {
		if authorized, err = sessionUsecase.SessionExists(c.Request().Context(), &models.Session{SessionId: cookie.Value}); err != nil {
			return false
		}
	}

	return authorized
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
			logger.LogrusLoggerWithContext(ctx.Request().Context()).Debug("Authorized!")

			return next(ctx)
		}
	}
}
