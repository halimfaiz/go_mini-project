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
	orderRepository := database.NewOrderRepository(db)
	paymentRepository := database.NewPaymentRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	productUsecase := usecase.NewProductUseCase(productRepository)
	cartUsecase := usecase.NewCartUseCase(cartRepository, productRepository)
	orderUsecase := usecase.NewOrderUsecase(orderRepository, cartRepository, userRepository, productRepository)
	paymentUsecase := usecase.NewPaymentUseCase(paymentRepository, userRepository, orderRepository)

	authController := controller.NewAuthController(userUsecase)
	userController := controller.NewUserController(userUsecase)
	productController := controller.NewProductController(productUsecase)
	cartController := controller.NewCartController(cartUsecase)
	orderController := controller.NewOrderController(orderUsecase)
	paymentController := controller.NewPaymentController(paymentUsecase, orderUsecase)

	e.Validator = &customValidator{validator: validator.New()}

	e.POST("/register/user", userController.RegisterUserController)
	e.POST("/register/admin", userController.RegisterAdminController)
	e.POST("/login/user", authController.LoginUserController)
	e.POST("/login/admin", authController.LoginAdminController)
	e.GET("/products", productController.GetListProductsController)

	user := e.Group("/user", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.GET("/carts", cartController.GetCartByUserID)
	user.POST("/carts/add", cartController.AddProductToCart)
	user.POST("/carts/remove", cartController.RemoveProductFromCart)
	user.POST("/orders", orderController.CreateOrderController)
	user.GET("/orders", orderController.GetOrderController)
	user.POST("/topup", userController.TopUpSaldoController)
	user.POST("/orders/checkout/:id", paymentController.CheckoutController)

	admin := e.Group("/admin", middleware.JWT([]byte(constant.SECRET_JWT)))
	admin.POST("/products", productController.AddProductController)
	admin.PUT("/products/:id", productController.UpdateProductController)
	admin.DELETE("/products/:id", productController.DeleteProductController)
	admin.GET("/orders/pending", orderController.GetOrderByStatusController)
	admin.GET("/orders", orderController.GetListOrderController)
	admin.PUT("/orders/:id", orderController.UpdateOrderStatusController)

}
