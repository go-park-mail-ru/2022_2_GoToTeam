package repositoryToUsecaseErrors

import "errors"

var ArticleRepositoryError = errors.New("error in article repository")

var ArticleRepositoryArticleDontExistsError = errors.New("error in article repository: article dont exists")
