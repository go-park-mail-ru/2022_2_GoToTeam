package models

type User struct {
	UserId   int
	Username string
	Email    string
	Login    string
	Password string
	//NickName   string `json:"nick_name"`
	FirstName  string
	LastName   string
	MiddleName string
}

type SignupData struct {
	UserName   string `json:"user_name"`
	FirstName  string
	LastName   string
	MiddleName string
	Email      string
	Login      string
}

type SignupResponse struct {
	Data    SignupData
	Message string
}
