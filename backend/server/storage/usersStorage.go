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

func (o *UsersStorage) AddUser(username string, email string, login string, password string) error {
	log.Println("Storage AddUser called.")

	user := &models.User{
		Username: username,
		Email:    email,
		Login:    login,
		Password: password,
	}

	o.mu.Lock()

	for _, v := range o.users {
		if v.Login == login {
			return errors.New("user with the same login exist")
		}
		if v.Email == email {
			return errors.New("user with the same email exist")
		}
	}

	user.UserId = o.nextID
	log.Println("New user id: ", user.UserId)
	o.nextID++
	o.users = append(o.users, user)

	o.mu.Unlock()

	return nil
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

	return nil, errors.New("user with the same login dont exists")
}

func (o *UsersStorage) GetUserByUid(uid int) (*models.User, error) {
	log.Println("Storage GetUserByLogin called.")

	o.mu.RLock()
	defer o.mu.RUnlock()

	for _, v := range o.users {
		if v.UserId == uid {
			return v, nil
		}
	}

	return nil, errors.New("user with the same login dont exists")
}

func (o *UsersStorage) GetUserByEmail(email string) (*models.User, error) {
	log.Println("Storage GetUserByLogin called.")

	o.mu.RLock()
	defer o.mu.RUnlock()

	for _, v := range o.users {
		if v.Email == email {
			return v, nil
		}
	}

	return nil, errors.New("user with the same email dont exists")
}
