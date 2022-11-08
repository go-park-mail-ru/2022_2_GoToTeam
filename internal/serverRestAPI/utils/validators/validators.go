package validators

import (
	"net/mail"
	"regexp"
	"unicode"
)

const (
	emailRFC5322RegExp = `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$` // RFC 5322. https://stackoverflow.com/questions/201323/how-to-validate-an-email-address-using-a-regular-expression/201378#201378  http://emailregex.com/  got for the javascript

	loginRegExp = `^[a-zA-Z][a-zA-Z0-9]{3,}$` // Starts from latin letter, after its possible latin letters and number, min length = 3

	passwordRegExp = `^[a-zA-Z0-9]{4,}$` // Only latins and numbers, min length = 4
)

func EmailIsValidByRegExp(email string) bool {
	match, err := regexp.MatchString(emailRFC5322RegExp, email)

	return err == nil && match
}

func EmailIsValidByCustomValidation(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func LoginIsValidByRegExp(login string) bool {
	match, err := regexp.MatchString(passwordRegExp, login)

	return err == nil && match
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

func PasswordIsValidByRegExp(password string) bool {
	match, err := regexp.MatchString(loginRegExp, password)

	return err == nil && match
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

func UsernameIsValidByCustomValidation(username string) bool {

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
