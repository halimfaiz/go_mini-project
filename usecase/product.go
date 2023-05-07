package usecase

import (
	"errors"
	"mini_project/model"
	"mini_project/repository/database"
)

type ProductUseCase interface {
	GetProducts() ([]model.Product, error)
	AddProduct(name, description string, price, stock uint) (*model.Product, error)
	DeleteProductById(id uint) error
	UpdateProduct(id, stock uint) (*model.Product, error)
}

type productUseCase struct {
	productRepository database.ProductRepository
}

func NewProductUseCase(productRepository database.ProductRepository) *productUseCase {
	return &productUseCase{productRepository: productRepository}
}

func (p *productUseCase) GetProducts() (products []model.Product, err error) {
	products, err = p.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productUseCase) AddProduct(name, description string, price, stock uint) (*model.Product, error) {
	products, err := p.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}
	for _, v := range products {
		if v.Name == name {
			return nil, errors.New("product already exist")
		}
	}
	product := &model.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	err = p.productRepository.AddProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productUseCase) UpdateProduct(id, stock uint) (*model.Product, error) {
	product, err := p.productRepository.GetProductById(id)
	if err != nil {
		return nil, err
	}
	if stock <= 0 {
		return nil, errors.New("stock can't be minus")
	} else {

		product.Stock = stock
		err = p.productRepository.UpdateProduct(product)
		if err != nil {
			return nil, err
		}
		return product, nil
	}
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
