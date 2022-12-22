package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type GetFeedError struct {
	Err error
}

func (gae *GetFeedError) Error() string {
	return gae.Err.Error()
}

type LoginIsNotValidError struct {
	Err error
}

func (linve *LoginIsNotValidError) Error() string {
	return linve.Err.Error()
}

type LoginDoesntExistError struct {
	Err error
}

func (ldee *LoginDoesntExistError) Error() string {
	return ldee.Err.Error()
}

type CategoryDoesntExistError struct {
	Err error
}

func (cdee *CategoryDoesntExistError) Error() string {
	return cdee.Err.Error()
}
