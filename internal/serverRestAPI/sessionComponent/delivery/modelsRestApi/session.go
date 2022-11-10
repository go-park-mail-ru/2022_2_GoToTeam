package modelsRestApi

type SessionInfo struct {
	Username      string `json:"username"`
	AvatarImgPath string `json:"avatar_img_path"`
}

type SessionCreate struct {
	UserData UserData `json:"user_data"`
}

type UserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
