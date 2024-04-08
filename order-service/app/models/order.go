package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pesto_coding/product_service/app/models"
	"time"
)

type Order struct {
	gorm.Model
	UserID          uint
	CreatedAt       time.Time
	OrderProductIds []models.Product `gorm:"foreignKey:OrderID"`
}

func (o *Order) TableName() string {
	return "orders"
}
