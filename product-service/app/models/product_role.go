package models

import "github.com/jinzhu/gorm"

type ProductRole struct {
	gorm.Model
	ProductID uint
	RoleID    uint
}
