package api

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

var ErrUserNotExist = Response{
	Status:  http.StatusFailedDependency,
	Message: "User doesn't exist",
}

var ErrUserAuthorised = Response{
	Status:  http.StatusFailedDependency,
	Message: "Already authorised",
}

var ErrUserNotAuthorised = Response{
	Status:  http.StatusFailedDependency,
	Message: "User not authorised",
}

var ErrAlreadyLogout = Response{
	Status:  http.StatusFailedDependency,
	Message: "Already logout",
}

var ErrNoNextFeedId = Response{
	Status:  http.StatusNotFound,
	Message: "Can`t found next articles",
}

var ErrWrongPassword = Response{
	Status:  http.StatusForbidden,
	Message: "Incorrect password",
}

var LogoutResponse = Response{
	Status:  http.StatusOK,
	Message: "Goodbye, we are waiting for you again",
}
