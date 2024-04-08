package models

import (
	"github.com/jinzhu/gorm"
)

type OrderStore struct {
	db *gorm.DB
}

func NewOrderStore(db *gorm.DB) *OrderStore {
	return &OrderStore{
		db: db,
	}
}

func (os *OrderStore) CreateOrder(order *Order) error {
	if err := os.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (os *OrderStore) UpdateOrderStatus(orderID uint, status string) error {
	order := &Order{}
	if err := os.db.First(order, orderID).Error; err != nil {
		return err
	}

	order.Status = status

	if err := os.db.Save(order).Error; err != nil {
		return err
	}

	return nil
}

func (os *OrderStore) GetOrderById(orderID uint) (*Order, error) {
	order := &Order{}
	if err := os.db.First(order, orderID).Error; err != nil {
		return nil, err
	}
	return order, nil
}
