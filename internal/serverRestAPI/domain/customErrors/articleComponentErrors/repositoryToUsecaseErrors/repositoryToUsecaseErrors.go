package repositoryToUsecaseErrors

import "errors"

var ArticleRepositoryError = errors.New("error in article repository")

var ArticleRepositoryArticleDoesntExistError = errors.New("error in article repository: article doesnt exist")
