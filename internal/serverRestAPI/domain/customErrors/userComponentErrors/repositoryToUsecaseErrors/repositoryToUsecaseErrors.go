package repositoryToUsecaseErrors

import "errors"

var UserRepositoryError = errors.New("error in user repository")

var UserRepositoryEmailDontExistsError = errors.New("error in user repository: email dont exists")

var UserRepositoryEmailExistsError = errors.New("error in user repository: email exists")

var UserRepositoryLoginDontExistsError = errors.New("error in user repository: login dont exists")

var UserRepositoryLoginExistsError = errors.New("error in user repository: login exists")
