package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/userComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/sessionUtils/httpCookieUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userUsecase    userComponentInterfaces.UserUsecaseInterface
	sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface
	logger         *logger.Logger
}

func NewUserController(userUsecase userComponentInterfaces.UserUsecaseInterface, sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface, logger *logger.Logger) *UserController {
	logger.LogrusLogger.Debug("Enter to the NewUserController function.")

	userController := &UserController{
		userUsecase:    userUsecase,
		sessionUsecase: sessionUsecase,
		logger:         logger,
	}

	logger.LogrusLogger.Info("UserController has created.")

	return userController
}

func (uc *UserController) SignupUserHandler(c echo.Context) error {
	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the SignupUserHandler function.")

	defer c.Request().Body.Close()

	parsedInput := new(modelsRestApi.User)
	if err := c.Bind(parsedInput); err != nil {
		uc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed input user json data: %#v", parsedInput)

	if err := uc.userUsecase.AddNewUser(c.Request().Context(), parsedInput.NewUserData.Email, parsedInput.NewUserData.Login, parsedInput.NewUserData.Username, parsedInput.NewUserData.Password); err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailIsNotValidError:
			return c.JSON(http.StatusBadRequest, "email is not valid")
		case *usecaseToDeliveryErrors.LoginIsNotValidError:
			return c.JSON(http.StatusBadRequest, "login is not valid")
		case *usecaseToDeliveryErrors.PasswordIsNotValidError:
			return c.JSON(http.StatusBadRequest, "password is not valid")
		case *usecaseToDeliveryErrors.EmailExistsError:
			return c.JSON(http.StatusConflict, "email exists")
		case *usecaseToDeliveryErrors.LoginExistsError:
			return c.JSON(http.StatusConflict, "login exists")
		default:
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	session, err := uc.sessionUsecase.CreateSessionForUser(c.Request().Context(), parsedInput.NewUserData.Email, parsedInput.NewUserData.Password)
	if err != nil {
		uc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	c.SetCookie(httpCookieUtils.MakeHttpCookie(session.SessionId))

	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Infof("User with the email %#v registered successful!", parsedInput.NewUserData.Email)

	return c.NoContent(http.StatusOK)
}

func (uc *UserController) UserInfoHandler(c echo.Context) error {
	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the UserInfoHandler function.")

	login := c.QueryParam("login")
	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed login: %#v", login)
	if login == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := uc.userUsecase.GetUserInfo(c.Request().Context(), login)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.LoginDoesntExistError:
			uc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusNotFound)
		case *usecaseToDeliveryErrors.LoginIsNotValidError:
			uc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusBadRequest) // TODO
		default:
			uc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	isSubscribed, err := uc.userUsecase.IsUserSubscribedOnUser(c.Request().Context(), login)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			uc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		default:
			uc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	userInfo := modelsRestApi.UserInfo{
		Username:         user.Username,
		RegistrationDate: user.RegistrationDate,
		SubscribersCount: user.SubscribersCount,
		Subscribed:       isSubscribed,
	}

	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed userInfo: ", userInfo)

	return c.JSON(http.StatusOK, userInfo)
}

func (uc *UserController) SubscribeHandler(c echo.Context) error {
	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the SubscribeHandler function.")

	defer c.Request().Body.Close()

	parsedInput := new(modelsRestApi.Subscribe)
	if err := c.Bind(parsedInput); err != nil {
		uc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed input json data: %#v", parsedInput)

	if err := uc.userUsecase.SubscribeOnUser(c.Request().Context(), parsedInput.Login); err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			uc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusInternalServerError)
		default:
			uc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Info("User subscribed successfully!")

	return c.NoContent(http.StatusOK)
}

func (uc *UserController) UnsubscribeHandler(c echo.Context) error {
	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the UnsubscribeHandler function.")

	defer c.Request().Body.Close()

	parsedInput := new(modelsRestApi.Subscribe)
	if err := c.Bind(parsedInput); err != nil {
		uc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed input json data: %#v", parsedInput)

	if err := uc.userUsecase.UnsubscribeFromUser(c.Request().Context(), parsedInput.Login); err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			uc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusInternalServerError)
		default:
			uc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	uc.logger.LogrusLoggerWithContext(c.Request().Context()).Info("User unsubscribed successfully!")

	return c.NoContent(http.StatusOK)
}
