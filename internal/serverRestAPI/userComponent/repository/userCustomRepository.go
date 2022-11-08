package repository

/*

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/modelsOLD"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"errors"
	"log"
	"sync"
)

var usersData = []*modelsOLD.User{
	{
		1,
		"asd",
		"asd@asd.asd",
		"asdLogin",
		"asdPass",
	},
	{
		2,
		"qwe",
		"qwe@qwe.qwe",
		"qweLogin",
		"qwePass",
	},
}

type usersStorage struct {
	users  []*modelsOLD.User
	mu     sync.RWMutex
	nextID int
	logger *logger.Logger
}

func NewUserCustomRepository(logger *logger.Logger) userComponentInterfaces.UserRepositoryInterface {
	return &usersStorage{
		users:  usersData,
		mu:     sync.RWMutex{},
		nextID: 3,
		logger: logger,
	}
}

func (us *usersStorage) PrintUsers() {
	log.Printf("Users in storage:")
	for _, v := range us.users {
		log.Printf("%#v ", v)
	}
}

func (us *usersStorage) AddUser(user *modelsOLD.User) error { // user_id
	log.Println("Storage AddUser called.")

	us.mu.Lock()
	defer us.mu.Unlock()

	for _, v := range us.users {
		if v.Login == user.Login {
			// TODO logger
			return errors.New("user with the same login exists")
		}
		if v.Email == user.Email {
			// TODO logger
			return errors.New("user with the same email exists")
		}
	}

	user.UserId = us.getIdForInsert()
	log.Println("New user id: ", user.UserId)
	us.users = append(us.users, user)
	us.PrintUsers()

	return nil
}

func (us *usersStorage) UserExistsByLogin(login string) bool {
	user, _ := us.GetUserByLogin(login)

	return user != nil
}

func (us *usersStorage) UserExistsByEmail(email string) bool {
	user, _ := us.GetUserByEmail(email)

	return user != nil
}

func (us *usersStorage) GetUserByLogin(login string) (*modelsOLD.User, error) {
	log.Println("Storage GetUserByLogin called.")

	us.mu.RLock()
	defer us.mu.RUnlock()

	for _, v := range us.users {
		if v.Login == login {
			return v, nil
		}
	}

	return nil, errors.New("user with this login dont exists")
}

func (us *usersStorage) GetUserByEmail(email string) (*modelsOLD.User, error) {
	log.Println("Storage GetUserByEmail called.")

	us.mu.RLock()
	defer us.mu.RUnlock()

	for _, v := range us.users {
		if v.Email == email {
			return v, nil
		}
	}

	return nil, errors.New("user with this email dont exists")
}

func (us *usersStorage) getIdForInsert() (id int) {
	// Deadlock:
	// us.mu.Lock()
	// defer us.mu.Unlock()

	id = us.nextID
	us.nextID++

	return
}
*/
