package app

import "github.com/jonascavalcanti/square-backend/internal/model"

type UserApp struct {
	repo UserRepository
}

func NewUserApp(repo UserRepository) *UserApp {
	return &UserApp{
		repo: repo,
	}
}

func (u *UserApp) CreateUser(user *model.User) error {
	// Business logic if needed
	return u.repo.CreateUser(user)
}

func (u *UserApp) GetUserByID(id uint) (*model.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UserApp) GetAllUsers() ([]*model.User, error) {
	return u.repo.GetAllUsers()
}

func (u *UserApp) UpdateUser(user *model.User) error {
	// Business logic if needed
	return u.repo.UpdateUser(user)
}

func (u *UserApp) DeleteUser(id uint) error {
	// Business logic if needed
	return u.repo.DeleteUser(id)
}
