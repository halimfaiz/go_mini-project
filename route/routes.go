package route

import (
	"mini_project/constant"
	"mini_project/controller"
	"mini_project/repository/database"
	"mini_project/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := database.NewUserRepository(db)
	productRepository := database.NewProductRepository(db)
	cartRepository := database.NewCartRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	productUsecase := usecase.NewProductUseCase(productRepository)
	cartUsecase := usecase.NewCartUseCase(cartRepository, productRepository)

	authController := controller.NewAuthController(userUsecase)
	userController := controller.NewUserController(userUsecase, userRepository)
	productController := controller.NewProductController(productUsecase, productRepository)
	cartController := controller.NewCartController(cartUsecase)

	e.Validator = &customValidator{validator: validator.New()}

	e.POST("/register/user", userController.RegisterUserController)
	e.POST("/register/admin", userController.RegisterAdminController)
	e.POST("/login/user", authController.LoginUserController)
	e.POST("/login/admin", authController.LoginAdminController)
	e.GET("/products", productController.GetListProducts)

	user := e.Group("/user", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.GET("/carts", cartController.GetCartByUserID)
	user.POST("/carts/add", cartController.AddProductToCart)
	user.POST("/carts/remove", cartController.RemoveProductFromCart)

	admin := e.Group("/admin", middleware.JWT([]byte(constant.SECRET_JWT)))
	admin.POST("/products", productController.AddProduct)
	admin.DELETE("/products/:id", productController.DeleteProduct)

}
