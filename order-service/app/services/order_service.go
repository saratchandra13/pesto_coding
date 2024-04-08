package services

import (
	"github.com/pesto_coding/order_service/app/models"
)

type OrderService struct {
	orderStore *models.OrderStore
}

func NewOrderService(orderStore *models.OrderStore) *OrderService {
	return &OrderService{
		orderStore: orderStore,
	}
}

func (os *OrderService) CreateOrder(userID uint, productIDs []uint) (*models.Order, error) {
	order := &models.Order{
		UserID: userID,
	}

	if err := os.orderStore.CreateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (os *OrderService) UpdateOrderStatus(orderID uint, status string) (*models.Order, error) {
	if err := os.orderStore.UpdateOrderStatus(orderID, status); err != nil {
		return nil, err
	}

	order, err := os.orderStore.GetOrderById(orderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
