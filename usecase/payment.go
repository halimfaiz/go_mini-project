package usecase

import (
	"errors"
	"mini_project/model"
	"mini_project/repository/database"
	"time"
)

type PaymentUseCase interface {
	Checkout(userID, orderID uint) (*model.Payment, error)
}

type paymentUseCase struct {
	paymentRepository database.PaymentRepository
	userRepository    database.UserRepository
	orderRepository   database.OrderRepository
}

func NewPaymentUseCase(
	paymentRepository database.PaymentRepository,
	userRepository database.UserRepository,
	orderRepository database.OrderRepository,
) *paymentUseCase {
	return &paymentUseCase{
		paymentRepository: paymentRepository,
		userRepository:    userRepository,
		orderRepository:   orderRepository,
	}
}

func (p *paymentUseCase) Checkout(userID, orderID uint) (payment *model.Payment, err error) {
	user, err := p.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	order, err := p.orderRepository.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	payment = &model.Payment{
		OrderID:     order.ID,
		TotalPrice:  order.TotalPrice,
		PaymentDate: time.Now(),
	}

	if order.Status == "UNPAID" {
		if user.Saldo < payment.TotalPrice {
			return nil, errors.New("insufficient balance")
		} else {
			user.Saldo -= payment.TotalPrice
			err = p.userRepository.UpdateUser(user)
			if err != nil {
				return nil, err
			}
		}
		err = p.paymentRepository.Checkout(payment)
		if err != nil {
			return nil, err
		}
		order.Status = "PAID"
		err = p.orderRepository.UpdateOrder(order)
		if err != nil {
			return nil, err
		}
		return payment, nil
	} else {
		return nil, errors.New("this order already paid")
	}
}
