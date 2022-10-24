package storage

import (
	"2022_2_GoTo_team/server/storage/models"
	"errors"
	"log"
	"sync"
)

type UsersStorage struct {
	users  []*models.User
	mu     sync.RWMutex
	nextID int
}

func GetUsersStorage() *UsersStorage {
	return &UsersStorage{
		users:  usersData,
		mu:     sync.RWMutex{},
		nextID: 3,
	}
}

func (o *UsersStorage) PrintUsers() {
	log.Printf("Users in storage:")
	for _, v := range o.users {
		log.Printf("%#v ", v)
	}
}

func (o *UsersStorage) AddUser(user *models.User) error { // user_id
	log.Println("Storage AddUser called.")

	o.mu.Lock()
	defer o.mu.Unlock()

	for _, v := range o.users {
		if v.Login == user.Login {
			return errors.New("user with the same login exist")
		}
		if v.Email == user.Email {
			return errors.New("user with the same email exist")
		}
	}

	user.UserId = o.getIdForInsert()
	log.Println("New user id: ", user.UserId)
	o.users = append(o.users, user)
	o.PrintUsers()

	return nil
}

func (o *UsersStorage) UserIsExistByLogin(login string) bool {
	user, _ := o.GetUserByLogin(login)

	return user != nil
}

func (o *UsersStorage) UserIsExistByEmail(email string) bool {
	user, _ := o.GetUserByEmail(email)

	return user != nil
}

func (o *UsersStorage) GetUserByLogin(login string) (*models.User, error) {
	log.Println("Storage GetUserByLogin called.")

	o.mu.RLock()
	defer o.mu.RUnlock()

	for _, v := range o.users {
		if v.Login == login {
			return v, nil
		}
	}

	return nil, errors.New("user with this login dont exists")
}

func (o *UsersStorage) GetUserByEmail(email string) (*models.User, error) {
	log.Println("Storage GetUserByEmail called.")

	o.mu.RLock()
	defer o.mu.RUnlock()

	for _, v := range o.users {
		if v.Email == email {
			return v, nil
		}
	}

	return nil, errors.New("user with this email dont exists")
}

func (o *UsersStorage) CreateUserInstanceFromData(username string, email string, login string, password string) *models.User {
	return &models.User{
		Username: username,
		Email:    email,
		Login:    login,
		Password: password,
	}
}

func (o *UsersStorage) getIdForInsert() (id int) {
	// Deadlock:
	// o.mu.Lock()
	// defer o.mu.Unlock()

	id = o.nextID
	o.nextID++

	return
}
