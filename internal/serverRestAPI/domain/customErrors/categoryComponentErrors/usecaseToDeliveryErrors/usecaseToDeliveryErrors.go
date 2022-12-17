package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type CategoryNotFoundError struct {
	Err error
}

func (cnfe *CategoryNotFoundError) Error() string {
	return cnfe.Err.Error()
}

type EmailForSessionDoesntExistError struct {
	Err error
}

func (efsdee *EmailForSessionDoesntExistError) Error() string {
	return efsdee.Err.Error()
}
