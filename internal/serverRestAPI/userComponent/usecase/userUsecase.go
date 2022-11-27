package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/errorsUtils"
	"2022_2_GoTo_team/pkg/logger"
	"2022_2_GoTo_team/pkg/validators"
	"context"
	"errors"
)

type userUsecase struct {
	userRepository userComponentInterfaces.UserRepositoryInterface
	logger         *logger.Logger
}

func NewUserUsecase(userRepository userComponentInterfaces.UserRepositoryInterface, logger *logger.Logger) userComponentInterfaces.UserUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewUserUsecase function.")

	userUsecase := &userUsecase{
		userRepository: userRepository,
		logger:         logger,
	}

	logger.LogrusLogger.Info("userUsecase has created.")

	return userUsecase
}

func (uu *userUsecase) GetUserInfo(ctx context.Context, login string) (*models.User, error) {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UserInfo function.")

	wrappingErrorMessage := "error while getting user info"

	if !validators.LoginIsValidByRegExp(login) {
		uu.logger.LogrusLoggerWithContext(ctx).Debugf("Login %s is not valid.", login)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.LoginIsNotValidError{Err: errors.New("login is not valid")})
	}

	user, err := uu.userRepository.GetUserInfo(ctx, login)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors.UserRepositoryLoginDoesntExistError:
			uu.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.LoginDoesntExistError{
				Err: err,
			})
		default:
			uu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{
				Err: err,
			})
		}
	}

	return user, nil
}

func (uu *userUsecase) AddNewUser(ctx context.Context, email string, login string, username string, password string) error {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the AddNewUser function.")

	wrappingErrorMessage := "error while adding the user"

	if err := uu.validateUserData(ctx, email, login, password); err != nil {
		uu.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}
	if err := uu.checkUserExists(ctx, email, login); err != nil {
		uu.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}
	if _, err := uu.userRepository.AddUser(ctx, email, login, username, password); err != nil {
		uu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return nil
}

func (uu *userUsecase) createUserInstanceFromData(ctx context.Context, email string, login string, username string, password string) *models.User {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the createUserInstanceFromData function.")

	return &models.User{
		Email:    email,
		Login:    login,
		Username: username,
		Password: password,
	}
}

func (uu *userUsecase) checkUserExists(ctx context.Context, email string, login string) error {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the checkUserExists function.")

	exists, err := uu.userExistsByEmail(ctx, email)
	if err != nil {
		uu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return &usecaseToDeliveryErrors.RepositoryError{Err: err}
	}
	if exists {
		uu.logger.LogrusLoggerWithContext(ctx).Debugf("User with this email %s exists.", email)
		return &usecaseToDeliveryErrors.EmailExistsError{Err: errors.New("user with this email exists")}
	}

	exists, err = uu.userExistsByLogin(ctx, login)
	if err != nil {
		uu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return &usecaseToDeliveryErrors.RepositoryError{Err: err}
	}
	if exists {
		uu.logger.LogrusLoggerWithContext(ctx).Debugf("User with this login %s exists.", login)
		return &usecaseToDeliveryErrors.LoginExistsError{Err: errors.New("user with this login exists")}
	}

	return nil
}

func (uu *userUsecase) userExistsByEmail(ctx context.Context, email string) (bool, error) {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the userExistsByEmail function.")

	exists, err := uu.userRepository.UserExistsByEmail(ctx, email)
	if err != nil {
		uu.logger.LogrusLoggerWithContext(ctx).Debug(err)
		return true, err
	}

	return exists, nil
}

func (uu *userUsecase) userExistsByLogin(ctx context.Context, login string) (bool, error) {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the userExistsByLogin function.")

	exists, err := uu.userRepository.UserExistsByLogin(ctx, login)
	if err != nil {
		uu.logger.LogrusLoggerWithContext(ctx).Debug(err)
		return true, err
	}

	return exists, nil
}

func (uu *userUsecase) validateUserData(ctx context.Context, email string, login string, password string) error {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the validateUserData function.")

	if !validators.EmailIsValidByCustomValidation(email) {
		uu.logger.LogrusLoggerWithContext(ctx).Debugf("Email %s is not valid.", email)
		return &usecaseToDeliveryErrors.EmailIsNotValidError{Err: errors.New("email is not valid")}
	}
	if !validators.LoginIsValidByRegExp(login) {
		uu.logger.LogrusLoggerWithContext(ctx).Debugf("Login %s is not valid.", login)
		return &usecaseToDeliveryErrors.LoginIsNotValidError{Err: errors.New("login is not valid")}
	}
	if !validators.PasswordIsValidByRegExp(password) {
		uu.logger.LogrusLoggerWithContext(ctx).Debug("Password is not valid.")
		return &usecaseToDeliveryErrors.PasswordIsNotValidError{Err: errors.New("password is not valid")}
	}

	return nil
}
