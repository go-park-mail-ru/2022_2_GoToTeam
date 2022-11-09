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
