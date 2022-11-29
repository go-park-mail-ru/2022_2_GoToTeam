package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/sessionComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfacesOLD"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/sessionComponentOLD/delivery/modelsRestApi"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/sessionUtils/httpCookieUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SessionController struct {
	sessionUsecase sessionComponentInterfacesOLD.SessionUsecaseInterface
	logger         *logger.Logger
}

func NewSessionController(sessionUsecase sessionComponentInterfacesOLD.SessionUsecaseInterface, logger *logger.Logger) *SessionController {
	logger.LogrusLogger.Debug("Enter to the NewSessionController function.")

	sessionController := &SessionController{
		sessionUsecase: sessionUsecase,
		logger:         logger,
	}

	logger.LogrusLogger.Info("SessionController has created.")

	return sessionController
}

func (sc *SessionController) CreateSessionHandler(c echo.Context) error {
	sc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the CreateSessionHandler function.")

	defer c.Request().Body.Close()
	parsedInput := new(modelsRestApi.SessionCreate)
	if err := c.Bind(parsedInput); err != nil {
		sc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	sc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed input: %#v", parsedInput)

	email := parsedInput.UserData.Email
	password := parsedInput.UserData.Password

	session, err := sc.sessionUsecase.CreateSessionForUser(c.Request().Context(), email, password)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailIsNotValidError:
			sc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.JSON(http.StatusBadRequest, "email is not valid")
		case *usecaseToDeliveryErrors.PasswordIsNotValidError:
			sc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.JSON(http.StatusBadRequest, "password is not valid")
		case *usecaseToDeliveryErrors.IncorrectEmailOrPasswordError:
			sc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.JSON(http.StatusBadRequest, "incorrect email or password")
		default:
			sc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	c.SetCookie(httpCookieUtils.MakeHttpCookie(session.SessionId))

	sc.logger.LogrusLoggerWithContext(c.Request().Context()).Infof("User email %#v auth success!", email)

	return c.NoContent(http.StatusOK)
}

func (sc *SessionController) RemoveSessionHandler(c echo.Context) error {
	sc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the RemoveSessionHandler function.")

	cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME)
	if err != nil {
		sc.logger.LogrusLoggerWithContext(c.Request().Context()).Info(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	if err := sc.sessionUsecase.RemoveSession(c.Request().Context(), &models.Session{SessionId: cookie.Value}); err != nil {
		sc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	httpCookieUtils.ExpireHttpCookie(cookie)
	c.SetCookie(cookie) // Need to reset new expired cookie

	sc.logger.LogrusLoggerWithContext(c.Request().Context()).Infof("User session %#v removed successfully.", cookie.Value)

	return c.NoContent(http.StatusOK)
}

func (sc *SessionController) SessionInfoHandler(c echo.Context) error {
	sc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the SessionInfoHandler function.")

	cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME)
	if err != nil {
		sc.logger.LogrusLoggerWithContext(c.Request().Context()).Info(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	user, err := sc.sessionUsecase.GetUserInfoBySession(c.Request().Context(), &models.Session{SessionId: cookie.Value})
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			sc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusNotFound)
		case *usecaseToDeliveryErrors.UserForSessionDoesntExistError:
			sc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			sc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	userInfoBySession := modelsRestApi.UserInfoBySession{
		Username:      user.Username,
		AvatarImgPath: user.AvatarImgPath,
	}
	sc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed userInfoBySession = ", userInfoBySession)

	return c.JSON(http.StatusOK, userInfoBySession)
}
