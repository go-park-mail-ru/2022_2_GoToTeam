package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
	"errors"
	"net/mail"
	"unicode"
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

	logger.LogrusLogger.Info("NewUserUsecase has created.")

	return userUsecase
}

func (uu *userUsecase) GetUserInfo(ctx context.Context, login string) (*models.User, error) {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UserInfo function.")

	wrappingErrorMessage := "error while getting user info"

	if !uu.loginIsValid(ctx, login) {
		uu.logger.LogrusLoggerWithContext(ctx).Debugf("Login %s is not valid.", login)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.LoginIsNotValidError{Err: errors.New("login is not valid")})
	}

	user, err := uu.userRepository.GetUserInfo(ctx, login)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors.UserRepositoryLoginDontExistsError:
			uu.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.LoginDontExistsError{
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

	if err := uu.validateUserData(ctx, email, login, username, password); err != nil {
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

func (uu *userUsecase) validateUserData(ctx context.Context, email string, login string, username string, password string) error {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the validateUserData function.")

	if !uu.emailIsValid(ctx, email) {
		uu.logger.LogrusLoggerWithContext(ctx).Debugf("Email %s is not valid.", email)
		return &usecaseToDeliveryErrors.EmailIsNotValidError{Err: errors.New("email is not valid")}
	}
	if !uu.loginIsValid(ctx, login) {
		uu.logger.LogrusLoggerWithContext(ctx).Debugf("Login %s is not valid.", login)
		return &usecaseToDeliveryErrors.LoginIsNotValidError{Err: errors.New("login is not valid")}
	}
	if !uu.usernameIsValid(ctx, username) {
		uu.logger.LogrusLoggerWithContext(ctx).Debugf("Username %s is not valid.", username)
		return &usecaseToDeliveryErrors.UsernameIsNotValidError{Err: errors.New("username is not valid")}
	}
	if !uu.passwordIsValid(ctx, password) {
		uu.logger.LogrusLoggerWithContext(ctx).Debug("Password is not valid.")
		return &usecaseToDeliveryErrors.PasswordIsNotValidError{Err: errors.New("password is not valid")}
	}

	return nil
}

func (uu *userUsecase) emailIsValid(ctx context.Context, email string) bool {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the emailIsValid function.")

	_, err := mail.ParseAddress(email)

	return err == nil
}

func (uu *userUsecase) loginIsValid(ctx context.Context, login string) bool {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the loginIsValid function.")

	if len(login) < 4 {
		return false
	}
	for _, sep := range login {
		if !unicode.IsLetter(sep) && sep != '_' {
			return false
		}
	}

	return true
}

// at least 8 symbols
// at least 1 upper symbol
// at least 1 special symbol
func (uu *userUsecase) passwordIsValid(ctx context.Context, password string) bool {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the passwordIsValid function.")

	letters := 0
	upper := false
	special := false
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		}
	}
	eightOrMore := letters >= 8

	return eightOrMore && upper && special
}

func (uu *userUsecase) usernameIsValid(ctx context.Context, username string) bool {
	uu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the usernameIsValid function.")

	if len(username) == 0 {
		return false
	}
	en := unicode.Is(unicode.Latin, rune(username[0]))

	for _, sep := range username {
		if (en && !unicode.Is(unicode.Latin, sep)) || (!en && unicode.Is(unicode.Latin, sep)) {
			return false
		}
	}

	return true
}
