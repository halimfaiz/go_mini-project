package usecase

import (
	"mini_project/middleware"
	"mini_project/model"
	"mini_project/repository/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	LoginUser(email, password, role string) (*model.User, error)
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepo database.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) CreateUser(user *model.User) error {

	err := u.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) GetUserByEmail(email string) (*model.User, error) {
	return u.userRepository.GetUserByEmail(email)
}

func (u *userUsecase) LoginUser(email, password, role string) (*model.User, error) {
	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "email not registered")
	}
	if user.Password != password {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid Password")
	}
	if user.Role != role {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	token, err := middleware.CreateToken(user.ID, user.Role)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "error create token")
	}
	user.Token = token
	return user, nil
}
