package repositoryToUsecaseErrors

import "errors"

var SessionRepositoryError = errors.New("error in session repository")

var SessionRepositoryEmailDoesntExistError = errors.New("error in session repository: email doesnt exist")
