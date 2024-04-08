package models

type OrderProduct struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  int
}

func (op *OrderProduct) TableName() string {
	return "order_products"
}
