package route

import (
	"mini_project/controller"
	"mini_project/repository/database"
	"mini_project/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
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

	userUsecase := usecase.NewUserUsecase(userRepository)

	authController := controller.NewAuthController(userUsecase)
	userController := controller.NewUserController(userUsecase, userRepository)

	e.Validator = &customValidator{validator: validator.New()}

	e.POST("/register/user", userController.RegisterUserController)
	e.POST("/register/admin", userController.RegisterAdminController)
	e.POST("/login/user", authController.LoginUserController)
	e.POST("/login/admin", authController.LoginAdminController)

}
