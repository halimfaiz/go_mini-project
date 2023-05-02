package database

import (
	"errors"
	"mini_project/config"
	"mini_project/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	LoginUser(email, password string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user *model.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) LoginUser(email, password string) error {
	var user model.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	if user.Password != password {
		return errors.New("wrong password")
	}
	return nil
}
