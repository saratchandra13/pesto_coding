package models

import "github.com/jinzhu/gorm"

type UserRole struct {
	gorm.Model
	UserId uint
	RoleId uint
}
