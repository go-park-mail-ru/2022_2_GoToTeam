package authSessionServiceErrors

import "errors"

var EmailIsNotValidError = errors.New("email is not valid")

var PasswordIsNotValidError = errors.New("password is not valid")

var IncorrectEmailOrPasswordError = errors.New("incorrect email or password")
