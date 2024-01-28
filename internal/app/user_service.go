package app

import "github.com/jonascavalcanti/square-backend/internal/model"

type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}
