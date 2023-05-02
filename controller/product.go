package controller

import (
	"mini_project/middleware"
	"mini_project/model"
	"mini_project/repository/database"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController interface {
	GetListProducts(c echo.Context) error
	AddProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

type productController struct {
	productUsecase    usecase.ProductUseCase
	productRepository database.ProductRepository
}

func NewProductController(
	productUsecase usecase.ProductUseCase,
	productRepository database.ProductRepository,
) *productController {
	return &productController{
		productUsecase,
		productRepository,
	}
}

func (p *productController) GetListProducts(c echo.Context) error {
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

func (p *productController) AddProduct(c echo.Context) error {
	_, err := middleware.ExtractTokenAdmin(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only admin can access this feature",
			"error":   err.Error(),
		})
	}

	var product model.Product

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}
	if err := p.productUsecase.AddProduct(&product); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add product",
		"product": product,
	})
}

func (p *productController) DeleteProduct(c echo.Context) error {
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
