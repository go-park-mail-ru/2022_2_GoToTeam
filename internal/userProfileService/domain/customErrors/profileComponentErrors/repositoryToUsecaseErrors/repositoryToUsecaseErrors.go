package repositoryToUsecaseErrors

import "errors"

var ProfileRepositoryError = errors.New("error in profile repository")

var ProfileRepositoryEmailDoesntExistError = errors.New("error in profile repository: email doesnt exist")

var ProfileRepositoryEmailExistsError = errors.New("error in profile repository: email exists")

var ProfileRepositoryLoginExistsError = errors.New("error in profile repository: login exists")
