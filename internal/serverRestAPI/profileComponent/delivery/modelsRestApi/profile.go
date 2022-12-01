package modelsRestApi

type Profile struct {
	Email         string `json:"email"`
	Login         string `json:"login"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	AvatarImgPath string `json:"avatar_img_path"`
}
