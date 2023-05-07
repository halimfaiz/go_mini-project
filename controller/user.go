package controller

import (
	"mini_project/middleware"
	"mini_project/model/payload"
	"mini_project/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	RegisterUserController(c echo.Context) error
	RegisterAdminController(c echo.Context) error
	TopUpSaldoController(c echo.Context) error
}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{userUsecase: userUsecase}
}

func (u *userController) RegisterUserController(c echo.Context) error {
	req := payload.CreateUserRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	if _, err := u.userUsecase.GetUserByEmail(req.Email); err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Email already registered",
		})
	}

	user, err := u.userUsecase.CreateUser(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to register user",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (u *userController) RegisterAdminController(c echo.Context) error {
	req := payload.CreateUserRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}
	if _, err := u.userUsecase.GetUserByEmail(req.Email); err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Email already registered",
		})
	}

	admin, err := u.userUsecase.CreateAdmin(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to register admin",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
		"user":    admin,
	})
}

func (u *userController) TopUpSaldoController(c echo.Context) error {
	userID, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"message": "only users can access this feature",
			"error":   err.Error(),
		})
	}
	payload := payload.TopUpRequest{}
	c.Bind(&payload)

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	user, err := u.userUsecase.TopUp(userID, payload.Saldo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Top Up Saldo",
		"user":    user,
	})
}
