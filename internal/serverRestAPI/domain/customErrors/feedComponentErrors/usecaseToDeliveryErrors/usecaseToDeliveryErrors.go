package usecaseToDeliveryErrors

type GetFeedError struct {
	Err error
}

func (gae *GetFeedError) Error() string {
	return gae.Err.Error()
}

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type LoginIsNotValidError struct {
	Err error
}

func (linve *LoginIsNotValidError) Error() string {
	return linve.Err.Error()
}

type LoginDontExistsError struct {
	Err error
}

func (ldee *LoginDontExistsError) Error() string {
	return ldee.Err.Error()
}

type CategoryDontExistsError struct {
	Err error
}

func (cdee *CategoryDontExistsError) Error() string {
	return cdee.Err.Error()
}
