package services

import (
	"github.com/lazycoder1995/pesto_coding/app/models"
)

type UserService struct {
	UserStore *models.UserStore
}

func NewUserService(userStore *models.UserStore) *UserService {
	return &UserService{
		UserStore: userStore,
	}
}

func (us *UserService) CreateUser(id string, username string, email string, password string) (*models.User, error) {
	user, err := models.NewUser(id, username, email, password)
	if err != nil {
		return nil, err
	}

	us.UserStore.AddUser(user)

	return user, nil
}

func (us *UserService) GetUser(id string) *models.User {
	return us.UserStore.GetUser(id)
}

func (us *UserService) AuthenticateUser(id string, password string) bool {
	user := us.UserStore.GetUser(id)
	if user == nil {
		return false
	}

	return user.CheckPassword(password)
}
