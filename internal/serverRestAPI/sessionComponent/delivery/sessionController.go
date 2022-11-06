package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/httpCookieUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type SessionController struct {
	sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface
	logger         *logger.Logger
}

func NewSessionController(sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface, logger *logger.Logger) *SessionController {
	return &SessionController{
		sessionUsecase: sessionUsecase,
		logger:         logger,
	}
}

func (sc *SessionController) isAuthorized(c echo.Context) bool {
	authorized := false
	if cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME); err == nil && cookie != nil {
		if authorized, err = sc.sessionUsecase.SessionExists(c.Request().Context(), &models.Session{SessionId: cookie.Value}); err != nil {
			return false
		}
	}

	return authorized
}

func (sc *SessionController) CreateSessionHandler(c echo.Context) error {
	defer c.Request().Body.Close()
	parsedInput := new(modelsRestApi.SessionCreate)
	if err := c.Bind(parsedInput); err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		sc.logger.LogrusLogger.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	//log.Println("parsedInput = ", parsedInput)
	sc.logger.LogrusLogger.Info("parsedInput = ", parsedInput)

	email := parsedInput.UserData.Email
	password := parsedInput.UserData.Password
	log.Println("URL", c.Request().URL)
	log.Println("email", email)
	log.Println("password ", password)

	session, err := sc.sessionUsecase.CreateSessionForUser(c.Request().Context(), email, password)
	if err != nil {
		// TODO logger
		log.Println("err in session controller: " + err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	c.SetCookie(httpCookieUtils.MakeHttpCookie(session.SessionId))

	sc.logger.LogrusLogger.Info("User auth success!")

	return c.NoContent(http.StatusOK)
}

func (sc *SessionController) RemoveSessionHandler(c echo.Context) error {
	if !sc.isAuthorized(c) {
		//c.LogrusLogger().Printf("Error: %s", "unauthorized")
		sc.logger.LogrusLogger.Error("unauthorized")
		return c.NoContent(http.StatusUnauthorized)
	}
	cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME)
	if err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		sc.logger.LogrusLogger.Error(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	if err := sc.sessionUsecase.RemoveSession(c.Request().Context(), &models.Session{SessionId: cookie.Value}); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	httpCookieUtils.ExpireHttpCookie(cookie)
	c.SetCookie(cookie) // Need to reset new expired cookie

	sc.logger.LogrusLogger.Info("User logout success")

	return c.NoContent(http.StatusOK)
}

func (sc *SessionController) SessionInfoHandler(c echo.Context) error {
	if !sc.isAuthorized(c) {
		//c.LogrusLogger().Printf("Error: %s", "unauthorized")
		sc.logger.LogrusLogger.Error("unauthorized")
		return c.NoContent(http.StatusUnauthorized)
	}
	cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME)
	if err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		sc.logger.LogrusLogger.Error(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	user, err := sc.sessionUsecase.GetUserBySession(c.Request().Context(), &models.Session{SessionId: cookie.Value})
	if err != nil {
		// TODO logger
		//api.logger.Error(err.Error())
		return c.NoContent(http.StatusUnauthorized)
	}

	sessionInfo := modelsRestApi.SessionInfo{
		Username: user.Username,
	}
	sc.logger.LogrusLogger.Info("Formed sessionInfo = ", sessionInfo)

	return c.JSON(http.StatusOK, sessionInfo)
}
