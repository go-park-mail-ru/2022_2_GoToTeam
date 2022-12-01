package repositoryToUsecaseErrors

import "errors"

var CategoryRepositoryError = errors.New("error in category repository")

var CategoryRepositoryCategoryDoesntExistError = errors.New("error in category repository: category doesnt exist")
