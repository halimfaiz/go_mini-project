package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       uint   `json:"price" form:"price"`
}
