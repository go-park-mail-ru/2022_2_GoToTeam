package restApiModels

type SessionInfo struct {
	Username string `json:"username"`
}

type SessionCreate struct {
	UserData UserData `json:"user_data"`
}

type UserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
