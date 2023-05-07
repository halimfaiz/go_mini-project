package database

import (
	"mini_project/config"
	"mini_project/model"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Checkout(payment *model.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db}
}

func (p *paymentRepository) Checkout(payment *model.Payment) error {
	err := config.DB.Create(payment).Error
	if err != nil {
		return err
	}
	return nil
}
