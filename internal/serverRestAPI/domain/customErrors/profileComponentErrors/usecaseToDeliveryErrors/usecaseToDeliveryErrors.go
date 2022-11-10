package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type EmailForSessionDontFoundError struct {
	Err error
}

func (efsdfe *EmailForSessionDontFoundError) Error() string {
	return efsdfe.Err.Error()
}

type UserForSessionDontFoundError struct {
	Err error
}

func (ufsdfe *UserForSessionDontFoundError) Error() string {
	return ufsdfe.Err.Error()
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
