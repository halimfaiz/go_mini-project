package usecase

import (
	"errors"
	"mini_project/model"
	"mini_project/repository/database"
)

type ProductUseCase interface {
	GetProducts() ([]model.Product, error)
	AddProduct(product *model.Product) error
	DeleteProductById(id uint) error
}

type productUseCase struct {
	productRepository database.ProductRepository
}

func NewProductUseCase(productRepository database.ProductRepository) *productUseCase {
	return &productUseCase{
		productRepository: productRepository,
	}
}

func (p *productUseCase) GetProducts() (products []model.Product, err error) {
	products, err = p.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productUseCase) AddProduct(product *model.Product) error {
	err := p.productRepository.AddProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productUseCase) DeleteProductById(id uint) error {
	product, err := p.productRepository.GetProductById(id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	err = p.productRepository.DeleteProductById(product)
	if err != nil {
		return err
	}
	return nil
}
