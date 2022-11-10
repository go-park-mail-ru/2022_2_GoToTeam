package repositoryToUsecaseErrors

import "errors"

var CategoryRepositoryError = errors.New("error in category repository")

var CategoryRepositoryCategoryDontExistsError = errors.New("error in category repository: category dont exists")
