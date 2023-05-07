package controller

import (
	"mini_project/middleware"
	"mini_project/model/payload"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController interface {
	CreateOrderController(c echo.Context) error
	GetOrderController(c echo.Context) error
	GetListOrderController(c echo.Context) error
	GetOrderByStatusController(c echo.Context) error
	UpdateOrderStatusController(c echo.Context) error
}

type orderController struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderController(orderUsecase usecase.OrderUsecase) *orderController {
	return &orderController{orderUsecase: orderUsecase}
}

func (o *orderController) CreateOrderController(c echo.Context) error {
	userId, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only users can access this feature",
			"error":   err.Error(),
		})
	}
	req := payload.DeliveryInfo{}
	c.Bind(&req)

	order, err := o.orderUsecase.CreateOrder(userId, req.Name, req.Address, req.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create order",
		"order":   order,
	})
}

func (o *orderController) GetOrderController(c echo.Context) error {
	userId, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only users can access this feature",
			"error":   err.Error(),
		})
	}

	orders, err := o.orderUsecase.GetOrderByUserID(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get User Orders",
		"order":   orders,
	})
}

func (o *orderController) GetOrderByStatusController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdmin(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access this feature",
			"error":   err.Error(),
		})
	}

	orders, err := o.orderUsecase.GetAllPaidOrder()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All Paid Orders",
		"orders":  orders,
	})
}

func (o *orderController) GetListOrderController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdmin(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access this feature",
			"error":   err.Error(),
		})
	}

	orders, err := o.orderUsecase.GetAllOrders()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All Orders",
		"orders":  orders,
	})
}

func (o *orderController) UpdateOrderStatusController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdmin(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access this feature",
			"error":   err.Error(),
		})
	}

	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Order ID",
			"err":     err.Error(),
		})
	}

	req := payload.StatusRequest{}
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Request Payload",
		})
	}
	order, err := o.orderUsecase.UpdateOrderStatus(uint(orderID), req.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update status order",
		"orders":  order,
	})
}
