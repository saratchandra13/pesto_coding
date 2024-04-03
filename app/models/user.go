package models

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string
	Username       string
	PhoneNumber    string
	Email          string
	Password       string `json:"password" binding:"required"` // this will not be stored in our database.
	Salt           string
	HashedPassword []byte
}

// NewUser creates a new User and hashes their password
func NewUser(id string, username string, email string, password string) (*User, error) {
	salt := generateSalt()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(salt+password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:             id,
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
