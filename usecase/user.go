package usecase

import (
	"mini_project/middleware"
	"mini_project/model"
	"mini_project/model/payload"
	"mini_project/repository/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	CreateUser(req *payload.CreateUserRequest) (*model.User, error)
	CreateAdmin(req *payload.CreateUserRequest) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	LoginUser(email, password, role string) (*model.User, error)
	TopUp(userID, saldo uint) (*model.User, error)
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepo database.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) CreateUser(req *payload.CreateUserRequest) (*model.User, error) {

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Phone:    req.Phone,
		Role:     "USER",
	}

	err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) CreateAdmin(req *payload.CreateUserRequest) (*model.User, error) {

	admin := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Phone:    req.Phone,
		Role:     "ADMIN",
	}

	err := u.userRepository.CreateUser(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (u *userUsecase) GetUserByEmail(email string) (*model.User, error) {
	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
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

func (u *userUsecase) TopUp(userID, saldo uint) (*model.User, error) {
	user, err := u.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	user.Saldo += saldo
	err = u.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
