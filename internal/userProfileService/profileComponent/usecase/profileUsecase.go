package usecase

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/customErrors/profileComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/userProfileService/domain/customErrors/profileComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/userProfileService/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/userProfileService/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"2022_2_GoTo_team/pkg/utils/validators"
	"context"
	"errors"
)

type profileUsecase struct {
	profileRepository profileComponentInterfaces.ProfileRepositoryInterface
	sessionRepository sessionComponentInterfaces.SessionRepositoryInterface
	logger            *logger.Logger
}

func NewProfileUsecase(profileRepository profileComponentInterfaces.ProfileRepositoryInterface, sessionRepository sessionComponentInterfaces.SessionRepositoryInterface, logger *logger.Logger) profileComponentInterfaces.ProfileUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewProfileUsecase function.")

	profileUsecase := &profileUsecase{
		profileRepository: profileRepository,
		sessionRepository: sessionRepository,
		logger:            logger,
	}

	logger.LogrusLogger.Info("profileUsecase has created.")

	return profileUsecase
}

func (pu *profileUsecase) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	pu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetProfileByEmail function.")

	wrappingErrorMessage := "error while getting profile by email"

	profile, err := pu.profileRepository.GetProfileByEmail(ctx, email)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors.ProfileRepositoryEmailDoesntExistError:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailDoesntExistError{Err: err})
		default:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return profile, nil
}

func (pu *profileUsecase) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	pu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateProfileBySession function.")

	wrappingErrorMessage := "error while updating newProfile by session"

	if err := pu.validateUserData(ctx, newProfile.Email, newProfile.Login, newProfile.Password); err != nil {
		pu.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	err := pu.profileRepository.UpdateProfileByEmail(ctx, newProfile, email)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors.ProfileRepositoryEmailExistsError:
			pu.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailExistsError{Err: err})
		case repositoryToUsecaseErrors.ProfileRepositoryLoginExistsError:
			pu.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.LoginExistsError{Err: err})
		default:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	// We should update sessions storage
	if newProfile.Email != email {
		if err := pu.sessionRepository.UpdateEmailBySession(ctx, session, newProfile.Email); err != nil {
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
		}
	}

	return nil
}

func (pu *profileUsecase) validateUserData(ctx context.Context, email string, login string, password string) error {
	pu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the validateUserData function.")

	if !validators.EmailIsValidByCustomValidation(email) {
		pu.logger.LogrusLoggerWithContext(ctx).Debugf("Email %s is not valid.", email)
		return &usecaseToDeliveryErrors.EmailIsNotValidError{Err: errors.New("email is not valid")}
	}
	if !validators.LoginIsValidByRegExp(login) {
		pu.logger.LogrusLoggerWithContext(ctx).Debugf("Login %s is not valid.", login)
		return &usecaseToDeliveryErrors.LoginIsNotValidError{Err: errors.New("login is not valid")}
	}
	if password != "" && !validators.PasswordIsValidByRegExp(password) {
		pu.logger.LogrusLoggerWithContext(ctx).Debug("Password is not valid.")
		return &usecaseToDeliveryErrors.PasswordIsNotValidError{Err: errors.New("password is not valid")}
	}

	return nil
}
