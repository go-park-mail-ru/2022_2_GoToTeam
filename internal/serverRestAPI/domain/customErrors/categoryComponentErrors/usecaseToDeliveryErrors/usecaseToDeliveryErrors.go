package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type CategoryDontExistsError struct {
	Err error
}

func (cdee *CategoryDontExistsError) Error() string {
	return cdee.Err.Error()
}
