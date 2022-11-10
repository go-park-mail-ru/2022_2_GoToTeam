package repositoryToUsecaseErrors

import "errors"

var SessionRepositoryError = errors.New("error in session repository")

var SessionRepositoryEmailDontExistsError = errors.New("error in session repository: email dont exists")
