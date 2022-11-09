package repositoryToUsecaseErrors

import "errors"

var ProfileRepositoryError = errors.New("error in profile repository")

var ProfileRepositoryEmailDontExistsError = errors.New("error in profile repository: email dont exists")

var ProfileRepositoryEmailExistsError = errors.New("error in profile repository: email exists")

var ProfileRepositoryLoginDontExistsError = errors.New("error in profile repository: login dont exists")

var ProfileRepositoryLoginExistsError = errors.New("error in profile repository: login exists")
