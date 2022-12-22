package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/profileComponent/delivery/modelsRestApi"
	domainPkg "2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
	"net/http"
)

type ProfileController struct {
	profileUsecase profileComponentInterfaces.ProfileUsecaseInterface
	logger         *logger.Logger
}

func NewProfileController(profileUsecase profileComponentInterfaces.ProfileUsecaseInterface, logger *logger.Logger) *ProfileController {
	logger.LogrusLogger.Debug("Enter to the NewProfileController function.")

	profileController := &ProfileController{
		profileUsecase: profileUsecase,
		logger:         logger,
	}

	logger.LogrusLogger.Info("ProfileController has created.")

	return profileController
}

func (pc *ProfileController) GetProfileHandler(c echo.Context) error {
	ctx := c.Request().Context()
	pc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetProfileHandler function.")

	email := ctx.Value(domainPkg.USER_EMAIL_KEY_FOR_CONTEXT).(string)
	pc.logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", email)
	if email == "" {
		pc.logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return c.NoContent(http.StatusUnauthorized)
	}

	profile, err := pc.profileUsecase.GetProfileByEmail(ctx, email)
	if err != nil {
		pc.logger.LogrusLoggerWithContext(ctx).Error(err)
		st, _ := status.FromError(err)
		return c.NoContent(errorsUtils.ExtractCodeFromGrpcErrorStatus(st))
	}

	profileOutput := modelsRestApi.Profile{
		Email:         profile.Email,
		Login:         profile.Login,
		Username:      profile.Username,
		AvatarImgPath: profile.AvatarImgPath,
	}
	pc.logger.LogrusLoggerWithContext(ctx).Debugf("Formed profileInfo: %#v", profileOutput)

	jsonBytes, err := profileOutput.MarshalJSON()
	if err != nil {
		pc.logger.LogrusLoggerWithContext(ctx).Error(err)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (pc *ProfileController) UpdateProfileHandler(c echo.Context) error {
	pc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the UpdateProfileHandler function.")
	defer c.Request().Body.Close()

	cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME)
	if err != nil {
		pc.logger.LogrusLoggerWithContext(c.Request().Context()).Info(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	parsedInputProfile := new(modelsRestApi.Profile)
	if err := c.Bind(parsedInputProfile); err != nil {
		pc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	pc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed parsedInputProfile: %#v", parsedInputProfile)

	ctx := c.Request().Context()
	email := ctx.Value(domainPkg.USER_EMAIL_KEY_FOR_CONTEXT).(string)
	pc.logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", email)
	if email == "" {
		pc.logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return c.NoContent(http.StatusUnauthorized)
	}

	err = pc.profileUsecase.UpdateProfileByEmail(c.Request().Context(), &models.Profile{Email: parsedInputProfile.Email, Login: parsedInputProfile.Login, Username: parsedInputProfile.Username, Password: parsedInputProfile.Password, AvatarImgPath: parsedInputProfile.AvatarImgPath}, email, &models.Session{SessionId: cookie.Value})
	if err != nil {
		pc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
		st, _ := status.FromError(err)
		return c.JSON(errorsUtils.ExtractCodeFromGrpcErrorStatus(st), st.Message())
	}

	return c.NoContent(http.StatusOK)
}
