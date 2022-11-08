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
