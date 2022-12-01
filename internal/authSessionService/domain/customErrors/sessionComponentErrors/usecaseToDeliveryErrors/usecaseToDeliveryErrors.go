package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type IncorrectEmailOrPasswordError struct {
	Err error
}

func (ieope *IncorrectEmailOrPasswordError) Error() string {
	return ieope.Err.Error()
}

type EmailForSessionDoesntExistError struct {
	Err error
}

func (efsdee *EmailForSessionDoesntExistError) Error() string {
	return efsdee.Err.Error()
}

type UserForSessionDoesntExistError struct {
	Err error
}

func (ufsdee *UserForSessionDoesntExistError) Error() string {
	return ufsdee.Err.Error()
}

type EmailIsNotValidError struct {
	Err error
}

func (einve *EmailIsNotValidError) Error() string {
	return einve.Err.Error()
}

type PasswordIsNotValidError struct {
	Err error
}

func (pinve *PasswordIsNotValidError) Error() string {
	return pinve.Err.Error()
}
