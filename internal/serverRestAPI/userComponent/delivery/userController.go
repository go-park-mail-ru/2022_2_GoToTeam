package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/userComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/httpCookieUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type UserController struct {
	userUsecase    userComponentInterfaces.UserUsecaseInterface
	sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface
	logger         *logger.Logger
}

func NewUserController(userUsecase userComponentInterfaces.UserUsecaseInterface, sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface, logger *logger.Logger) *UserController {
	userController := &UserController{
		userUsecase:    userUsecase,
		sessionUsecase: sessionUsecase,
		logger:         logger,
	}

	return userController
}

func (uc *UserController) SignupUserHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	parsedInput := new(modelsRestApi.User)
	if err := c.Bind(parsedInput); err != nil {
		uc.logger.LogrusLogger.Error(err)
		//c.Logger().Printf("Error: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	log.Println("Parsed input user data:", parsedInput)

	if err := uc.userUsecase.AddNewUser(parsedInput.NewUserData.Email, parsedInput.NewUserData.Login, parsedInput.NewUserData.Username, parsedInput.NewUserData.Password); err != nil {
		switch errors.Unwrap(err).(type) {
		case *userComponentErrors.EmailIsNotValidError:
			return c.NoContent(http.StatusBadRequest)
		case *userComponentErrors.LoginIsNotValidError:
			return c.NoContent(http.StatusBadRequest)
		case *userComponentErrors.UsernameIsNotValidError:
			return c.NoContent(http.StatusBadRequest)
		case *userComponentErrors.PasswordIsNotValidError:
			return c.NoContent(http.StatusBadRequest)
		case *userComponentErrors.EmailExistsError:
			return c.NoContent(http.StatusConflict)
		case *userComponentErrors.LoginExistsError:
			return c.NoContent(http.StatusConflict)
		default:
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	session, err := uc.sessionUsecase.CreateSessionForUser(parsedInput.NewUserData.Email, parsedInput.NewUserData.Password)
	if err != nil {
		// TODO logger
		log.Println("err: " + err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	c.SetCookie(httpCookieUtils.MakeHttpCookie(session.SessionId))

	uc.logger.LogrusLogger.Info("User register successful!")

	return c.NoContent(http.StatusOK)
}
