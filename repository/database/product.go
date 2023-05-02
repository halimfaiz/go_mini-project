package database

import (
	"mini_project/config"
	"mini_project/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductById(id uint) (*model.Product, error)
	GetProducts() ([]model.Product, error)
	AddProduct(product *model.Product) error
	DeleteProductById(product *model.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (p *productRepository) GetProductById(id uint) (*model.Product, error) {
	var product model.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product
	err := config.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepository) AddProduct(product *model.Product) error {
	err := config.DB.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) DeleteProductById(product *model.Product) error {
	err := config.DB.Delete(product).Error
	if err != nil {
		return err
	}
	return nil
}
