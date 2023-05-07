package database

import (
	"mini_project/config"
	"mini_project/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *model.Order) error
	CreateOrderItems(orderItem *model.OrderItem) error
	GetOrderByID(orderID uint) (*model.Order, error)
	GetOrderByUserID(userID uint) ([]*model.Order, error)
	GetAllOrders() ([]*model.Order, error)
	UpdateOrder(order *model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (o *orderRepository) CreateOrder(order *model.Order) error {
	if err := config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) CreateOrderItems(orderItem *model.OrderItem) error {
	if err := config.DB.Create(orderItem).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) GetOrderByID(orderID uint) (*model.Order, error) {
	var order model.Order
	if err := config.DB.Preload("OrderItems.Product").First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *orderRepository) GetOrderByUserID(userID uint) ([]*model.Order, error) {
	var order []*model.Order

	if err := config.DB.Model(&model.Order{}).Preload("OrderItems.Product").Where("user_id = ?", userID).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (o *orderRepository) GetAllOrders() ([]*model.Order, error) {
	var orders []*model.Order

	if err := config.DB.Preload("OrderItems.Product").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderRepository) UpdateOrder(order *model.Order) error {
	if err := config.DB.Save(order).Error; err != nil {
		return err
	}
	return nil
}
