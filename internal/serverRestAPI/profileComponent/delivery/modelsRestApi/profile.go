package modelsRestApi

type Profile struct {
	Email         string `json:"email"`
	Login         string `json:"login"`
	Username      string `json:"username"`
	AvatarImgPath string `json:"avatar_img_path"`
}
