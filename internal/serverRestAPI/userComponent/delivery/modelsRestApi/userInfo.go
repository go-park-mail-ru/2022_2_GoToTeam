package modelsRestApi

type UserInfo struct {
	Username         string `json:"username"`
	RegistrationDate string `json:"registration_date"`
	SubscribersCount int    `json:"subscribers_count"`
}
