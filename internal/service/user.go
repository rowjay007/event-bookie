package service

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/repository"
)

// UserService provides user-related services
type UserService struct {
	UserRepo *repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// CreateUser creates a new user
func (us *UserService) CreateUser(user *models.User) error {
	return us.UserRepo.CreateUser(user)
}

// GetUserByID retrieves a user by its ID
func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	return us.UserRepo.GetUserByID(id)
}

// UpdateUser updates an existing user
func (us *UserService) UpdateUser(user *models.User) error {
	return us.UserRepo.UpdateUser(user)
}

// DeleteUser deletes a user
func (us *UserService) DeleteUser(user *models.User) error {
	return us.UserRepo.DeleteUser(user)
}
