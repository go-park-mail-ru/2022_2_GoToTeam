package validators

import (
	"net/mail"
	"regexp"
	"unicode"
)

const (
	EMAIL_RFC5322_REGEXP_STRING = `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$` // RFC 5322. https://stackoverflow.com/questions/201323/how-to-validate-an-email-address-using-a-regular-expression/201378#201378  http://emailregex.com/  got for the javascript

	LOGIN_REGEXP_STRING = `^[a-zA-Z][a-zA-Z0-9]{3,}$` // Starts from latin letter, after its possible latin letters and number, min length = 4

	//PASSWORD_REGEXP_STRING = `^[a-zA-Z0-9]{4,}$` // Only latins and numbers, min length = 4
	PASSWORD_REGEXP_STRING = `^[a-zA-Z0-9!@#$%^&*]{4,}$`
)

var (
	emailRegExp = regexp.MustCompile(EMAIL_RFC5322_REGEXP_STRING)

	loginRegExp = regexp.MustCompile(LOGIN_REGEXP_STRING)

	passwordRegExp = regexp.MustCompile(PASSWORD_REGEXP_STRING)
)

func EmailIsValidByRegExp(email string) bool {
	return emailRegExp.MatchString(email)
}

func LoginIsValidByRegExp(login string) bool {
	return loginRegExp.MatchString(login)
}

func PasswordIsValidByRegExp(password string) bool {
	return passwordRegExp.MatchString(password)
}

func EmailIsValidByCustomValidation(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func LoginIsValidByCustomValidation(login string) bool {
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
func PasswordIsValidByCustomValidation(password string) bool {
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
