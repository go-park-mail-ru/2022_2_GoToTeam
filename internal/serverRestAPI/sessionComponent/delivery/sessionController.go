package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/delivery/restApiModels"
	"2022_2_GoTo_team/internal/utils/logger"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

const SESSION_HEADER_NAME = "session_id"

type SessionHandler struct {
	sessionUsecase interfaces.SessionUsecaseInterface
	logger         *logger.Logger
}

func NewSessionHandler(sessionUsecase interfaces.SessionUsecaseInterface, logger *logger.Logger) *SessionHandler {
	sessionHandler := &SessionHandler{
		sessionUsecase: sessionUsecase,
		logger:         logger,
	}

	return sessionHandler
}

func (sh *SessionHandler) isAuthorized(c echo.Context) bool {
	authorized := false
	if cookie, err := c.Cookie(SESSION_HEADER_NAME); err == nil && cookie != nil {
		authorized = sh.sessionUsecase.IsSessionExists(&models.Session{Cookie: cookie})
	}

	return authorized
}

func (sh *SessionHandler) CreateSessionHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	parsedInput := new(restApiModels.SessionCreate)
	if err := c.Bind(parsedInput); err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		sh.logger.LogrusLogger.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	//log.Println("parsedInput = ", parsedInput)
	sh.logger.LogrusLogger.Info("parsedInput = ", parsedInput)

	email := parsedInput.UserData.Email
	password := parsedInput.UserData.Password
	log.Println("URL", c.Request().URL)
	log.Println("email", email)
	log.Println("password ", password)

	session, err := sh.sessionUsecase.CreateSessionForUser(email, password, SESSION_HEADER_NAME)
	if err != nil {
		// TODO logger
		log.Println("err in session controller: " + err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	c.SetCookie(session.Cookie)

	sh.logger.LogrusLogger.Info("User auth success!")

	return c.NoContent(http.StatusOK)
}

func (sh *SessionHandler) RemoveSessionHandler(c echo.Context) error {
	if !sh.isAuthorized(c) {
		//c.LogrusLogger().Printf("Error: %s", "unauthorized")
		sh.logger.LogrusLogger.Error("unauthorized")
		return c.NoContent(http.StatusUnauthorized)
	}
	cookie, err := c.Cookie(SESSION_HEADER_NAME)
	if err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		sh.logger.LogrusLogger.Error(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	sh.sessionUsecase.RemoveSession(&models.Session{Cookie: cookie})
	c.SetCookie(cookie) // Need to reset new expired cookie

	sh.logger.LogrusLogger.Info("User logout success")
	return c.NoContent(http.StatusOK)
}

func (sh *SessionHandler) SessionInfoHandler(c echo.Context) error {
	if !sh.isAuthorized(c) {
		//c.LogrusLogger().Printf("Error: %s", "unauthorized")
		sh.logger.LogrusLogger.Error("unauthorized")
		return c.NoContent(http.StatusUnauthorized)
	}
	cookie, err := c.Cookie(SESSION_HEADER_NAME)
	if err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		sh.logger.LogrusLogger.Error(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	user, err := sh.sessionUsecase.GetUserBySession(&models.Session{Cookie: cookie})
	if err != nil {
		// TODO logger
		//api.logger.Error(err.Error())
		return c.NoContent(http.StatusUnauthorized)
	}

	sessionInfo := restApiModels.SessionInfo{
		Username: user.Username,
	}
	sh.logger.LogrusLogger.Info("Formed sessionInfo = ", sessionInfo)

	return c.JSON(http.StatusOK, sessionInfo)
}
