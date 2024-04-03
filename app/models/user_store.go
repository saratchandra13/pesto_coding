package models

import (
	"sync"
)

type UserStore struct {
	m     sync.RWMutex
	Users map[string]*User
}

func NewUserStore() *UserStore {
	return &UserStore{
		Users: make(map[string]*User),
	}
}

func (us *UserStore) AddUser(user *User) {
	us.m.Lock()
	defer us.m.Unlock()

	us.Users[user.ID] = user
}

func (us *UserStore) GetUser(id string) *User {
	us.m.RLock()
	defer us.m.RUnlock()

	return us.Users[id]
}
