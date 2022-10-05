package models

type Session struct {
	UserData UserData `json:"user_data"`
}

type UserData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
