package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type EmailForSessionNotFoundError struct {
	Err error
}

func (efsnfe *EmailForSessionNotFoundError) Error() string {
	return efsnfe.Err.Error()
}

type UserForSessionNotFoundError struct {
	Err error
}

func (ufsnfe *UserForSessionNotFoundError) Error() string {
	return ufsnfe.Err.Error()
}

type EmailExistsError struct {
	Err error
}

func (eee *EmailExistsError) Error() string {
	return eee.Err.Error()
}

type LoginExistsError struct {
	Err error
}

func (lee *LoginExistsError) Error() string {
	return lee.Err.Error()
}

type EmailIsNotValidError struct {
	Err error
}

func (einve *EmailIsNotValidError) Error() string {
	return einve.Err.Error()
}

type LoginIsNotValidError struct {
	Err error
}

func (linve *LoginIsNotValidError) Error() string {
	return linve.Err.Error()
}

type PasswordIsNotValidError struct {
	Err error
}

func (pinve *PasswordIsNotValidError) Error() string {
	return pinve.Err.Error()
}
