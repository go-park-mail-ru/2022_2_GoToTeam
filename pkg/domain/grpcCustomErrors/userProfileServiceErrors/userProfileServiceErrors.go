package userProfileServiceErrors

import "errors"

var EmailIsNotValidError = errors.New("email is not valid")

var LoginIsNotValidError = errors.New("login is not valid")

var PasswordIsNotValidError = errors.New("password is not valid")

var EmailExistsError = errors.New("email exists")

var LoginExistsError = errors.New("login exists")
