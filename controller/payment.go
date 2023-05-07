package controller

import (
	"mini_project/middleware"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController interface {
	GetListPaymentsController(c echo.Context) error
	AddPaymentController(c echo.Context) error
	DeletePaymentController(c echo.Context) error
}

type paymentController struct {
	paymentUsecase usecase.PaymentUseCase
	orderUsecase   usecase.OrderUsecase
}

func NewPaymentController(
	paymentUsecase usecase.PaymentUseCase,
	orderUsecase usecase.OrderUsecase,
) *paymentController {
	return &paymentController{
		paymentUsecase: paymentUsecase,
		orderUsecase:   orderUsecase,
	}
}

func (p *paymentController) CheckoutController(c echo.Context) error {
	userId, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only users can access this feature",
			"error":   err.Error(),
		})
	}

	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Order ID",
		})
	}

	order, err := p.orderUsecase.GetOrderByID(uint(orderID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	payment, err := p.paymentUsecase.Checkout(userId, order.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Checkout order",
		"orders":  payment,
	})
}
