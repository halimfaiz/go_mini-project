package middleware

import (
	"errors"
	"mini_project/constant"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}

func ExtractTokenAdmin(e echo.Context) (uint, error) {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "admin" {
		return 0, errors.New("not authorized")

	}
	return uint(claims["id"].(float64)), nil
}

func ExtractTokenUser(e echo.Context) (uint, error) {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "user" {
		return 0, errors.New("not authorized")

	}
	return uint(claims["id"].(float64)), nil
}
