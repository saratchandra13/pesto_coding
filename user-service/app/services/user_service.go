package services

import (
	"github.com/pesto_coding/user_service/app/models"
)

type UserService struct {
	UserStore *models.UserStore
}

func NewUserService(userStore *models.UserStore) *UserService {
	return &UserService{
		UserStore: userStore,
	}
}

func (us *UserService) CreateUser(username string, email string, password string) (*models.User, error) {
	// call the user store's NewUser method
	return us.UserStore.CreateUser(username, email, password)
}

func (us *UserService) GetUser(id uint) *models.User {
	// call the user store's GetUserById method
	return us.UserStore.GetUserById(id)
}

func (us *UserService) AuthenticateUser(username string, password string) bool {
	// call the user store's AuthenticateUser method
	return us.UserStore.AuthenticateUser(username, password)
}
