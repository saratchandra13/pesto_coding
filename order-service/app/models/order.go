package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID          uint
	OrderProductIds []OrderProduct `gorm:"foreignKey:OrderID"`
	Status          string         `gorm:"default:'pending'"`
}

func (o *Order) TableName() string {
	return "orders"
}
