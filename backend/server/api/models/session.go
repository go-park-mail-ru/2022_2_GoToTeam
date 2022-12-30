package models

type Session struct {
	UserData UserData `json:"user_data"`
}

type UserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
