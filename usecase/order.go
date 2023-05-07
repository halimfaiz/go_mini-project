package usecase

import (
	"errors"
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

type OrderUsecase interface {
	CreateOrder(userId uint, name, address, phone string) (order *model.Order, err error)
	GetOrderByID(orderID uint) (*model.Order, error)
	GetOrderByUserID(UserID uint) ([]*model.Order, error)
	GetAllOrders() ([]*model.Order, error)
	GetAllPaidOrder() ([]*model.Order, error)
	UpdateOrderStatus(orderID uint, status string) (*model.Order, error)
}

type orderUsecase struct {
	orderRepository   database.OrderRepository
	cartRepository    database.CartRepository
	userRepository    database.UserRepository
	productRepository database.ProductRepository
}

func NewOrderUsecase(
	orderRepository database.OrderRepository,
	cartRepository database.CartRepository,
	userRepository database.UserRepository,
	productRepository database.ProductRepository,
) *orderUsecase {
	return &orderUsecase{
		orderRepository:   orderRepository,
		cartRepository:    cartRepository,
		userRepository:    userRepository,
		productRepository: productRepository,
	}
}

func (o *orderUsecase) CreateOrder(userId uint, name, address, phone string) (order *model.Order, err error) {
	user, err := o.userRepository.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	if name != "" {
		user.Name = name
	}
	if address != "" {
		user.Address = address
	}
	if phone != "" {
		user.Phone = phone
	}

	cart, err := o.cartRepository.GetCartByUserID(user.ID)
	if err != nil {
		return nil, errors.New("cart not found")
	}

	order = &model.Order{
		UserID:     userId,
		Name:       user.Name,
		Address:    user.Address,
		Phone:      user.Phone,
		TotalPrice: cart.TotalPrice,
		Status:     "UNPAID",
	}

	if err := o.orderRepository.CreateOrder(order); err != nil {
		return nil, err
	}

	if err := o.cartRepository.DeleteCart(cart); err != nil {
		return nil, err
	}

	var orderItems []model.OrderItem
	for _, v := range cart.CartProducts {
		orderItem := &model.OrderItem{
			OrderID:   order.ID,
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
			Product:   v.Product,
		}
		o.orderRepository.CreateOrderItems(orderItem)
		orderItems = append(orderItems, *orderItem)
		product, err := o.productRepository.GetProductById(orderItem.ProductID)
		if err != nil {
			return nil, err
		}
		if product.Stock < v.Product.Stock {
			return nil, fmt.Errorf("cannot order %d product. theres only %d product left", v.Product.Stock, product.Stock)
		} else {
			product.Stock -= v.Product.Stock
			err = o.productRepository.UpdateProduct(product)
			if err != nil {
				return nil, err
			}
		}
	}
	order.OrderItems = orderItems
	err = o.orderRepository.UpdateOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *orderUsecase) GetOrderByID(orderID uint) (*model.Order, error) {
	orders, err := o.orderRepository.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderUsecase) GetOrderByUserID(UserID uint) ([]*model.Order, error) {

	orders, err := o.orderRepository.GetOrderByUserID(UserID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderUsecase) GetAllOrders() ([]*model.Order, error) {
	orders, err := o.orderRepository.GetAllOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderUsecase) GetAllPaidOrder() ([]*model.Order, error) {
	orders, err := o.orderRepository.GetAllOrders()
	if err != nil {
		return nil, err
	}
	paid := []*model.Order{}
	for _, order := range orders {
		if order.Status == "PAID" {
			paid = append(paid, order)
		}
	}
	return paid, nil
}

func (o *orderUsecase) UpdateOrderStatus(orderID uint, status string) (*model.Order, error) {
	order, err := o.orderRepository.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}
	order.Status = status
	err = o.orderRepository.UpdateOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil

}
