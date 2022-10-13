package models

import "net/http"

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

var ErrUnpackingJSON = ErrorResponse{
	Status:  http.StatusUnprocessableEntity,
	Message: "Error unpacking JSON",
}

var ErrUserExist = ErrorResponse{
	Status:  http.StatusFailedDependency,
	Message: "User is already exist",
}

var ErrUserAuthorised = ErrorResponse{
	Status:  http.StatusFailedDependency,
	Message: "Already authorised",
}
