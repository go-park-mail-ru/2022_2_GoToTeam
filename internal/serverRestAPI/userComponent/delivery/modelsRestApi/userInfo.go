package modelsRestApi

import "time"

type UserInfo struct {
	Username         string    `json:"username"`
	RegistrationDate time.Time `json:"registration_date"`
	SubscribersCount int       `json:"subscribers_count"`
}
