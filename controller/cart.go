package controller

import (
	"mini_project/middleware"
	"mini_project/model/payload"
	"mini_project/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartController interface {
	AddProductToCart(c echo.Context) error
	RemoveProductFromCart(c echo.Context) error
	GetCartByUserID(c echo.Context) error
}

type cartController struct {
	cartUsecase usecase.CartUseCase
}

func NewCartController(
	cartUsecase usecase.CartUseCase,
) *cartController {
	return &cartController{
		cartUsecase: cartUsecase,
	}
}

func (cc *cartController) AddProductToCart(c echo.Context) error {
	userId, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only users can access this feature",
			"error":   err.Error(),
		})
	}

	req := payload.AddProductToCartRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	err = cc.cartUsecase.AddProductToCart(userId, req.ProductID, req.Quantity)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success Add Product to Cart",
	})
}

func (cc *cartController) RemoveProductFromCart(c echo.Context) error {
	userId, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only users can access this feature",
			"error":   err.Error(),
		})
	}

	req := payload.AddProductToCartRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	err = cc.cartUsecase.RemoveProductFromCart(userId, req.ProductID, req.Quantity)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success Remove Product from Cart",
	})
}

func (cc *cartController) GetCartByUserID(c echo.Context) error {
	userID, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only users can access this feature",
			"error":   err.Error(),
		})
	}

	cart, err := cc.cartUsecase.GetCartByUserID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "Cart Not Found",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Cart",
		"Cart":    cart,
	})
}
