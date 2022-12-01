package repositoryToUsecaseErrors

import "errors"

var UserRepositoryError = errors.New("error in user repository")

var UserRepositoryEmailDoesntExistError = errors.New("error in user repository: email doesnt exist")

var UserRepositoryEmailExistsError = errors.New("error in user repository: email exists")

var UserRepositoryLoginDoesntExistError = errors.New("error in user repository: login doesnt exist")

var UserRepositoryLoginExistsError = errors.New("error in user repository: login exists")
