package controller

import (
	"errors"
	"mini_project/model"
	"mini_project/model/payload"
	"mini_project/repository/database"
	"mini_project/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	RegisterUserController(c echo.Context) error
	RegisterAdminController(c echo.Context) error
}

type userController struct {
	userUsecase    usecase.UserUsecase
	userRepository database.UserRepository
}

func NewUserController(
	userUsecase usecase.UserUsecase,
	userRepository database.UserRepository,
) *userController {
	return &userController{
		userUsecase,
		userRepository,
	}
}

func (u *userController) RegisterUserController(c echo.Context) error {
	req := payload.CreateUserRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	if _, err := u.userRepository.GetUserByEmail(req.Email); err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Email already registered",
		})
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Phone:    req.Phone,
		Role:     "user",
	}
	if err := u.userUsecase.CreateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to register user",
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

	if _, err := u.userRepository.GetUserByEmail(req.Email); err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Email already registered",
		})
	}

	admin := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Phone:    req.Phone,
		Role:     "admin",
	}
	if admin.Role != "admin" {
		return errors.New("invalid role")
	}
	if err := u.userUsecase.CreateUser(admin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to register admin",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
		"user":    admin,
	})
}
