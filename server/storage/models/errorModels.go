package models

import "net/http"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

var ErrUnpackingJSON = Response{
	Status:  http.StatusUnprocessableEntity,
	Message: "Error unpacking JSON",
}

var ErrUserExist = Response{
	Status:  http.StatusFailedDependency,
	Message: "User is already exist",
}

var ErrUserAuthorised = Response{
	Status:  http.StatusFailedDependency,
	Message: "Already authorised",
}

var ErrAlreadyLogout = Response{
	Status:  http.StatusFailedDependency,
	Message: "Already logout",
}

var LogoutResponse = Response{
	Status:  http.StatusOK,
	Message: "Goodbye, we are waiting for you again",
}
