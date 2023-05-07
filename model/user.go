package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Saldo    uint   `json:"saldo"`
	Token    string `gorm:"-"`
}
