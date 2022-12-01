package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type ArticleDoesntExistError struct {
	Err error
}

func (adee *ArticleDoesntExistError) Error() string {
	return adee.Err.Error()
}

type EmailForSessionDoesntExistError struct {
	Err error
}

func (efsdee *EmailForSessionDoesntExistError) Error() string {
	return efsdee.Err.Error()
}
