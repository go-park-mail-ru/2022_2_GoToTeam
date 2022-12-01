package delivery

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/customErrors/profileComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/userProfileService/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcCustomErrors/userProfileServiceErrors"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/userProfileServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"google.golang.org/grpc/status"
	"net/http"
)

type ProfileDelivery struct {
	userProfileServiceGrpcProtos.UnimplementedUserProfileServiceServer

	profileUsecase profileComponentInterfaces.ProfileUsecaseInterface
	logger         *logger.Logger
}

func NewProfileDelivery(profileUsecase profileComponentInterfaces.ProfileUsecaseInterface, logger *logger.Logger) *ProfileDelivery {
	logger.LogrusLogger.Debug("Enter to the NewProfileDelivery function.")

	profileDelivery := &ProfileDelivery{
		profileUsecase: profileUsecase,
		logger:         logger,
	}

	logger.LogrusLogger.Info("ProfileController has created.")

	return profileDelivery
}

func (pd *ProfileDelivery) GetProfileByEmail(ctx context.Context, email *userProfileServiceGrpcProtos.UserEmail) (*userProfileServiceGrpcProtos.Profile, error) {
	pd.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetProfileByEmail function.")
	pd.logger.LogrusLoggerWithContext(ctx).Debugf("Input email: %#v", email)

	profile, err := pd.profileUsecase.GetProfileByEmail(ctx, email.Email)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailDoesntExistError:
			pd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, status.Errorf(http.StatusUnauthorized, "")
		default:
			pd.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, status.Errorf(http.StatusInternalServerError, "")
		}
	}

	profileOutput := &userProfileServiceGrpcProtos.Profile{
		Email:         profile.Email,
		Login:         profile.Login,
		Password:      profile.Password,
		Username:      profile.Username,
		AvatarImgPath: profile.AvatarImgPath,
	}
	pd.logger.LogrusLoggerWithContext(ctx).Debugf("Formed profile = %#v, %#v, %#v, %#v, %#v",
		profileOutput.Email,
		profileOutput.Login,
		profileOutput.Password,
		profileOutput.Username,
		profileOutput.AvatarImgPath,
	)

	return profileOutput, nil
}

func (pd *ProfileDelivery) UpdateProfileByEmail(ctx context.Context, updateProfileData *userProfileServiceGrpcProtos.UpdateProfileData) (*userProfileServiceGrpcProtos.Nothing, error) {
	pd.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateProfileByEmail function.")
	pd.logger.LogrusLoggerWithContext(ctx).Debugf("Input updateProfileData: %#v", updateProfileData)

	parsedInputProfile := &models.Profile{
		Email:         updateProfileData.Profile.Email,
		Login:         updateProfileData.Profile.Login,
		Password:      updateProfileData.Profile.Password,
		Username:      updateProfileData.Profile.Username,
		AvatarImgPath: updateProfileData.Profile.AvatarImgPath,
	}

	pd.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed parsedInputProfile: %#v, input email: %#v, sessionId: %#v", parsedInputProfile, updateProfileData.Email, updateProfileData.SessionId)

	err := pd.profileUsecase.UpdateProfileByEmail(ctx, parsedInputProfile, updateProfileData.Email, &models.Session{SessionId: updateProfileData.SessionId})
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailIsNotValidError:
			pd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, status.Errorf(http.StatusBadRequest, userProfileServiceErrors.EmailIsNotValidError.Error())
		case *usecaseToDeliveryErrors.LoginIsNotValidError:
			pd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, status.Errorf(http.StatusBadRequest, userProfileServiceErrors.LoginIsNotValidError.Error())
		case *usecaseToDeliveryErrors.PasswordIsNotValidError:
			pd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, status.Errorf(http.StatusBadRequest, userProfileServiceErrors.PasswordIsNotValidError.Error())
		case *usecaseToDeliveryErrors.EmailExistsError:
			pd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, status.Errorf(http.StatusConflict, userProfileServiceErrors.EmailExistsError.Error())
		case *usecaseToDeliveryErrors.LoginExistsError:
			pd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, status.Errorf(http.StatusConflict, userProfileServiceErrors.LoginExistsError.Error())
		default:
			pd.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, status.Errorf(http.StatusInternalServerError, "")
		}
	}

	return &userProfileServiceGrpcProtos.Nothing{Ok: true}, nil
}
