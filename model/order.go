package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	Name       string      `json:"name"`
	Address    string      `json:"address"`
	Phone      string      `json:"phone"`
	TotalPrice uint        `json:"total_price"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Product   Product `json:"product"`
}
