package controller

import (
	"mini_project/middleware"
	"mini_project/model/payload"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController interface {
	GetListProductsController(c echo.Context) error
	AddProductController(c echo.Context) error
	DeleteProductController(c echo.Context) error
	UpdateProductController(c echo.Context) error
}

type productController struct {
	productUsecase usecase.ProductUseCase
}

func NewProductController(productUsecase usecase.ProductUseCase) *productController {
	return &productController{productUsecase: productUsecase}
}

func (p *productController) GetListProductsController(c echo.Context) error {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all products",
		"products": products,
	})
}

func (p *productController) AddProductController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdmin(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access this feature",
			"error":   err.Error(),
		})
	}

	req := payload.ProductRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	product, err := p.productUsecase.AddProduct(req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add product",
		"product": product,
	})
}

func (p *productController) DeleteProductController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdmin(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access this feature",
			"error":   err.Error(),
		})
	}

	productID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Product ID",
		})
	}

	if err := p.productUsecase.DeleteProductById(uint(productID)); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "Record Not Found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Product",
	})
}

func (p *productController) UpdateProductController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdmin(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access this feature",
			"error":   err.Error(),
		})
	}
	productID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Product ID",
		})
	}

	req := payload.UpdateStockRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	product, err := p.productUsecase.UpdateProduct(uint(productID), req.Stock)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Update stock product",
		"product": product,
	})
}
