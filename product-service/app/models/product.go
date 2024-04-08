package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`
	Category    string
	ImageURL    string
	Stock       int    `gorm:"not null"`
	Roles       []Role `gorm:"many2many:product_roles;"`
	Version     int    `gorm:"default:1"`
}
