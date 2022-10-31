package models

type User struct {
	NewUserData NewUserData `json:"new_user_data"`
}

type NewUserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
