package controller

import (
	"mini_project/model/payload"
	"mini_project/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	LoginUserController(c echo.Context) error
	LoginAdminController(c echo.Context) error
}

type authController struct {
	userUsecase usecase.UserUsecase
}

func NewAuthController(userUsecase usecase.UserUsecase) *authController {
	return &authController{userUsecase}
}

func (a *authController) LoginUserController(c echo.Context) error {
	payload := payload.LoginRequest{}
	c.Bind(&payload)

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	user, err := a.userUsecase.LoginUser(payload.Email, payload.Password, "USER")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    user,
	})
}

func (a *authController) LoginAdminController(c echo.Context) error {
	payload := payload.LoginRequest{}
	c.Bind(&payload)

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	user, err := a.userUsecase.LoginUser(payload.Email, payload.Password, "ADMIN")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    user,
	})
}
