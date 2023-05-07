package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID       uint          `json:"user_id"`
	TotalPrice   uint          `json:"total_price"`
	Status       string        `json:"status"`
	CartProducts []CartProduct `json:"cart_products"`
}

type CartProduct struct {
	gorm.Model
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Product   Product `json:"product"`
}
