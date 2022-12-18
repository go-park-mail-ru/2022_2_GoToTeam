package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type OpenFileError struct {
	Err error
}

func (ofe *OpenFileError) Error() string {
	return ofe.Err.Error()
}

type FileSizeBigError struct {
	Err error
}

func (fsbe *FileSizeBigError) Error() string {
	return fsbe.Err.Error()
}

type NotImageError struct {
	Err error
}

func (nie *NotImageError) Error() string {
	return nie.Err.Error()
}

type EmailForSessionDoesntExistError struct {
	Err error
}

func (efsdee *EmailForSessionDoesntExistError) Error() string {
	return efsdee.Err.Error()
}
