package repositoryToUsecaseErrors

import "errors"

var CommentaryRepositoryError = errors.New("error in commentary repository")

var CommentaryRepositoryCommentaryDoesntExistError = errors.New("error in commentary repository: article doesnt exist")
