package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"errors"
	"fmt"
	"net/mail"
	"unicode"
)

type userUsecase struct {
	userRepository userComponentInterfaces.UserRepositoryInterface
	logger         *logger.Logger
}

func NewUserUsecase(userRepository userComponentInterfaces.UserRepositoryInterface, logger *logger.Logger) userComponentInterfaces.UserUsecaseInterface {
	userUsecase := &userUsecase{
		userRepository: userRepository,
		logger:         logger,
	}
	// TODO logger
	userUsecase.userRepository.PrintUsers()

	return userUsecase
}

func (uu *userUsecase) AddNewUser(email string, login string, username string, password string) error {
	if err := uu.validateUserData(email, login, username, password); err != nil {
		return fmt.Errorf("can not add user: %w", err)
	}
	if err := uu.checkUserExists(email, login); err != nil {
		return fmt.Errorf("can not add user: %w", err)
	}
	if err := uu.userRepository.AddUser(uu.createUserInstanceFromData(username, email, login, password)); err != nil {
		return fmt.Errorf("can not add user: %w", err)
	}

	return nil
}

func (uu *userUsecase) createUserInstanceFromData(username string, email string, login string, password string) *models.User {
	return &models.User{
		Username: username,
		Email:    email,
		Login:    login,
		Password: password,
	}
}

func (uu *userUsecase) checkUserExists(email string, login string) error {
	if uu.userExistsByEmail(email) {
		uu.logger.LogrusLogger.Debugf("User with this email %s exists.", email)
		return &userComponentErrors.EmailExistsError{Err: errors.New("user with this email exists")}
	}
	if uu.userExistsByLogin(login) {
		uu.logger.LogrusLogger.Debugf("User with this login %s exists.", login)
		return &userComponentErrors.LoginExistsError{Err: errors.New("user with this login exists")}
	}

	return nil
}

func (uu *userUsecase) userExistsByEmail(email string) bool {
	return uu.userRepository.UserExistsByEmail(email)
}

func (uu *userUsecase) userExistsByLogin(login string) bool {
	return uu.userRepository.UserExistsByLogin(login)
}

func (uu *userUsecase) validateUserData(email string, login string, username string, password string) error {
	if !uu.emailIsValid(email) {
		uu.logger.LogrusLogger.Debugf("Email %s is not valid.", email)
		return &userComponentErrors.EmailIsNotValidError{Err: errors.New("email is not valid")}
	}
	if !uu.loginIsValid(login) {
		uu.logger.LogrusLogger.Debugf("Login %s is not valid.", login)
		return &userComponentErrors.LoginIsNotValidError{Err: errors.New("login is not valid")}
	}
	if !uu.usernameIsValid(username) {
		uu.logger.LogrusLogger.Debugf("Username %s is not valid.", username)
		return &userComponentErrors.UsernameIsNotValidError{Err: errors.New("username is not valid")}
	}
	if !uu.passwordIsValid(password) {
		uu.logger.LogrusLogger.Debug("Password is not valid.")
		return &userComponentErrors.PasswordIsNotValidError{Err: errors.New("password is not valid")}
	}

	return nil
}

func (uu *userUsecase) emailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func (uu *userUsecase) loginIsValid(login string) bool {
	if len(login) < 8 {
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAA")
		return false
	}
	for _, sep := range login {
		if !unicode.IsLetter(sep) && sep != '_' {
			fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBB")
			return false
		}
	}

	return true
}

// at least 8 symbols
// at least 1 upper symbol
// at least 1 special symbol
func (uu *userUsecase) passwordIsValid(password string) bool {
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

func (uu *userUsecase) usernameIsValid(username string) bool {
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
