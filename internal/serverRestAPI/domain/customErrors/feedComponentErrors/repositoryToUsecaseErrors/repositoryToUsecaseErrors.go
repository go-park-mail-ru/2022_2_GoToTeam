package repositoryToUsecaseErrors

import "errors"

var FeedRepositoryError = errors.New("error in feed repository")

var FeedRepositoryEmailDontExistsError = errors.New("error in feed repository: email dont exists")

var FeedRepositoryEmailExistsError = errors.New("error in feed repository: email exists")

var FeedRepositoryLoginDontExistsError = errors.New("error in feed repository: login dont exists")

var FeedRepositoryLoginExistsError = errors.New("error in feed repository: login exists")
