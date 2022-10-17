package models

type User struct {
	UserId   int
	Username string `json:"username"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	//NickName   string `json:"nick_name"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
}

type LoginForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignupData struct {
	UserName   string `json:"username"`
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
