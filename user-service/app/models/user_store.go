package models

import "github.com/jinzhu/gorm"

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) CreateUser(username string, email string, password string) (*User, error) {
	user, err := NewUser(username, email, password)
	if err != nil {
		return nil, err
	}

	return user, us.db.Create(user).Error
}

func (us *UserStore) GetUserById(id uint) *User {
	var user User
	us.db.Find(&user, id)
	return &user
}

func (us *UserStore) GetUserByUsername(username string) *User {
	var user User
	us.db.Find(&user, "username = ?", username)
	return &user
}

func (us *UserStore) AuthenticateUser(username string, password string) bool {
	user := us.GetUserByUsername(username)
	if user == nil {
		return false
	}

	return user.CheckPassword(password)
}
