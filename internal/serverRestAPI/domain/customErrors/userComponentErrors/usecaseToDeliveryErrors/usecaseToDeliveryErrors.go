package usecaseToDeliveryErrors

type AddUserError struct {
	Err error
}

func (aue *AddUserError) Error() string {
	return aue.Err.Error()
}

type GetUserInfoError struct {
	Err error
}

func (gui *GetUserInfoError) Error() string {
	return gui.Err.Error()
}

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type EmailExistsError struct {
	Err error
}

func (eee *EmailExistsError) Error() string {
	return eee.Err.Error()
}

type EmailDontExistsError struct {
	Err error
}

func (ede *EmailDontExistsError) Error() string {
	return ede.Err.Error()
}

type LoginExistsError struct {
	Err error
}

func (lee *LoginExistsError) Error() string {
	return lee.Err.Error()
}

type LoginDontExistsError struct {
	Err error
}

func (lde *LoginDontExistsError) Error() string {
	return lde.Err.Error()
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
