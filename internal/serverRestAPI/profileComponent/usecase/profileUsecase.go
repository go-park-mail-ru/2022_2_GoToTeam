package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/profileComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/profileComponentErrors/usecaseToDeliveryErrors"
	repositoryToUsecaseErrors_sessionComponent "2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/sessionComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/validators"
	"context"
	"errors"
)

type profileUsecase struct {
	profileRepository profileComponentInterfaces.ProfileRepositoryInterface
	sessionRepository sessionComponentInterfaces.SessionRepositoryInterface
	logger            *logger.Logger
}

func NewProfileUsecase(profileRepository profileComponentInterfaces.ProfileRepositoryInterface, sessionRepository sessionComponentInterfaces.SessionRepositoryInterface, logger *logger.Logger) profileComponentInterfaces.ProfileUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewCategoryUsecase function.")

	profileUsecase := &profileUsecase{
		profileRepository: profileRepository,
		sessionRepository: sessionRepository,
		logger:            logger,
	}

	logger.LogrusLogger.Info("profileUsecase has created.")

	return profileUsecase
}

func (pu *profileUsecase) GetProfileBySession(ctx context.Context, session *models.Session) (*models.Profile, error) {
	pu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetProfileBySession function.")

	wrappingErrorMessage := "error while getting profile by session"

	email, err := pu.sessionRepository.GetEmailBySession(ctx, session)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors_sessionComponent.SessionRepositoryEmailDontExistsError:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDontFoundError{Err: err})
		default:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	profile, err := pu.profileRepository.GetProfileByEmail(ctx, email)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors.ProfileRepositoryEmailDontExistsError:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.UserForSessionDontFoundError{Err: err})
		default:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return profile, nil
}

func (pu *profileUsecase) UpdateProfileBySession(ctx context.Context, newProfile *models.Profile, session *models.Session) error {
	pu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateProfileBySession function.")

	wrappingErrorMessage := "error while updating newProfile by session"

	if err := pu.validateUserData(ctx, newProfile.Email, newProfile.Login, newProfile.Password); err != nil {
		pu.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	email, err := pu.sessionRepository.GetEmailBySession(ctx, session)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors_sessionComponent.SessionRepositoryEmailDontExistsError:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDontFoundError{Err: err})
		default:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	err = pu.profileRepository.UpdateProfileByEmail(ctx, newProfile, email)
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
		pu.sessionRepository.UpdateEmailBySession(ctx, session, newProfile.Email)
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
	if !validators.PasswordIsValidByRegExp(password) {
		pu.logger.LogrusLoggerWithContext(ctx).Debug("Password is not valid.")
		return &usecaseToDeliveryErrors.PasswordIsNotValidError{Err: errors.New("password is not valid")}
	}

	return nil
}
