package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type TagDoesntExistError struct {
	Err error
}

func (tdee *TagDoesntExistError) Error() string {
	return tdee.Err.Error()
}
