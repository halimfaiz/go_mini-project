package database

import (
	"mini_project/config"
	"mini_project/model"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(cart *model.Cart) error
	CreateCartProduct(cart *model.CartProduct) error
	GetCartByUserID(userID uint) (*model.Cart, error)
	UpdateCart(cart *model.Cart) error
	UpdateCartProduct(cartProduct *model.CartProduct) error
	DeleteCartProduct(cartProduct *model.CartProduct) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db}
}

func (c *cartRepository) CreateCart(cart *model.Cart) error {
	err := config.DB.Create(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) CreateCartProduct(cartProduct *model.CartProduct) error {
	err := config.DB.Create(cartProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) GetCartByUserID(userID uint) (*model.Cart, error) {
	cart := &model.Cart{}
	err := config.DB.Model(&model.Cart{}).Preload("CartItems.Product").Where("user_id = ? AND status = ?", userID, "active").First(cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *cartRepository) UpdateCart(cart *model.Cart) error {
	err := config.DB.Save(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) UpdateCartProduct(cartProduct *model.CartProduct) error {
	err := config.DB.Save(cartProduct).Error
	if err != nil {
		return err
	}
	return nil
}
func (c *cartRepository) DeleteCartProduct(cartProduct *model.CartProduct) error {
	err := config.DB.Delete(cartProduct).Error
	if err != nil {
		return err
	}
	return nil
}
