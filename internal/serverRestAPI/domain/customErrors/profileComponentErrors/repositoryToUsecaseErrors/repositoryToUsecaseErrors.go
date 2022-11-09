package repositoryToUsecaseErrors

import "errors"

var ProfileRepositoryError = errors.New("error in profile repository")

var ProfileRepositoryEmailDontExistsError = errors.New("error in profile repository: email dont exists")
