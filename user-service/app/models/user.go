package models

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	PhoneNumber    string
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"-"`
	Salt           string
	HashedPassword []byte `gorm:"not null"`
	Roles          []Role `gorm:"many2many:user_roles;"`
}

// NewUser creates a new User and hashes their password
func NewUser(username string, email string, password string) (*User, error) {
	salt := generateSalt()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(salt+password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username:       username,
		Email:          email,
		Salt:           salt,
		HashedPassword: hashedPassword,
	}

	return user, nil
}

// CheckPassword checks if the provided password is correct
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(u.Salt+password))
	return err == nil
}

func generateSalt() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
