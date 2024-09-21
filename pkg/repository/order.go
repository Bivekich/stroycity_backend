package repository

import (
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order model.Order) error {
	order.Status = "Processing"
	return r.db.Create(&order).Error
}

func (r *OrderRepository) GetOrderById(orderID int) (model.Order, error) {
	var order model.Order
	err := r.db.Preload("OrderItems").Preload("OrderItems.Item").Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}
