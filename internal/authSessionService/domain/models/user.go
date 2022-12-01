package models

type User struct {
	UserId             int
	Email              string
	Login              string
	Password           string
	Username           string
	Sex                string
	DateOfBirth        string
	AvatarImgPath      string
	RegistrationDate   string
	SubscribersCount   int
	SubscriptionsCount int
}
