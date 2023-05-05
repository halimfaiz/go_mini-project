package usecase

import (
	"errors"
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

type CartUseCase interface {
	AddProductToCart(userID, productID, quantity uint) error
	RemoveProductFromCart(userID, productID, quantity uint) error
	GetCartByUserID(userID uint) (*model.Cart, error)
}

type cartUseCase struct {
	cartRepository    database.CartRepository
	productRepository database.ProductRepository
}

func NewCartUseCase(
	cartRepository database.CartRepository,
	productRepository database.ProductRepository,
) *cartUseCase {
	return &cartUseCase{
		cartRepository:    cartRepository,
		productRepository: productRepository,
	}
}

func (c *cartUseCase) AddProductToCart(userID, productID, quantity uint) error {
	product, err := c.productRepository.GetProductById(productID)
	if err != nil {
		return errors.New("product not found")
	}

	//check if user have cart.
	cart, err := c.cartRepository.GetCartByUserID(userID)
	if err != nil {
		// Jika user belum memiliki cart, maka buat cart baru dengan status "active"
		cart = &model.Cart{
			UserID: userID,
			Status: "active",
		}
		err = c.cartRepository.CreateCart(cart)
		if err != nil {
			return err
		}
	}

	//check if product already exist in cart
	var cartProduct *model.CartProduct
	for _, cp := range cart.CartItems {
		if cp.Product.ID == productID {
			cartProduct = &cp
			break
		}
	}
	//if product exist in cart
	if cartProduct != nil {
		cartProduct.Quantity += quantity
		err = c.cartRepository.UpdateCartProduct(cartProduct)
		if err != nil {
			return err
		}
	} else {
		cartProduct = &model.CartProduct{
			CartID:    cart.ID,
			ProductID: product.ID,
			Quantity:  quantity,
		}
		err = c.cartRepository.CreateCartProduct(cartProduct)
		if err != nil {
			return err
		}
	}
	cart.TotalPrice += (product.Price * quantity)
	err = c.cartRepository.UpdateCart(cart)
	if err != nil {
		return err
	}
	return nil
}

func (c *cartUseCase) RemoveProductFromCart(userID, productID, quantity uint) error {
	cart, err := c.cartRepository.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	//check if product already exist in cart
	var cartProduct *model.CartProduct
	for _, cp := range cart.CartItems {
		if cp.Product.ID == productID {
			cartProduct = &cp
			break
		}
	}

	if cartProduct == nil {
		return fmt.Errorf("product is not available in the cart")
	}

	if cartProduct.Quantity < quantity {
		return fmt.Errorf("cannot remove %d product. theres only %d product left in cart", quantity, cartProduct.Quantity)
	}

	if cartProduct.Quantity > quantity {
		cart.TotalPrice -= (cartProduct.Product.Price * quantity)
		cartProduct.Quantity -= quantity

		err = c.cartRepository.UpdateCart(cart)
		if err != nil {
			return err
		}

		err = c.cartRepository.UpdateCartProduct(cartProduct)
		if err != nil {
			return err
		}
	} else {
		cart.TotalPrice -= (cartProduct.Product.Price * cartProduct.Quantity)

		err = c.cartRepository.UpdateCart(cart)
		if err != nil {
			return err
		}

		err = c.cartRepository.UpdateCartProduct(cartProduct)
		if err != nil {
			return err
		}

		err = c.cartRepository.DeleteCartProduct(cartProduct)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *cartUseCase) GetCartByUserID(userID uint) (*model.Cart, error) {
	cart, err := c.cartRepository.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
