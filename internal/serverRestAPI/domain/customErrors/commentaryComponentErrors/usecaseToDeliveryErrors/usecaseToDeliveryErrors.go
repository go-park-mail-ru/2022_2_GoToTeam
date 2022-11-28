package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type EmailForSessionDoesntExistError struct {
	Err error
}

func (efsdee *EmailForSessionDoesntExistError) Error() string {
	return efsdee.Err.Error()
}
