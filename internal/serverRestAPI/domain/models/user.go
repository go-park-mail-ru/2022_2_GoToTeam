package models

import "time"

type User struct {
	UserId             int
	Email              string
	Login              string
	Password           string
	Username           string
	Sex                string
	DateOfBirth        time.Time
	AvatarImgPath      string
	RegistrationDate   time.Time
	SubscribersCount   int
	SubscriptionsCount int
}
