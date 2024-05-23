package service

import (
    "github.com/rowjay007/event-bookie/internal/models"
    "github.com/rowjay007/event-bookie/internal/repository"
)

// UserService provides user-related services
type UserService struct {
    UserRepository *repository.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(userRepo *repository.UserRepository) *UserService {
    return &UserService{UserRepository: userRepo}
}

// CreateUser creates a new user
func (us *UserService) CreateUser(user *models.User) error {
    return us.UserRepository.Create(user)
}

// GetAllUsers retrieves all users
func (us *UserService) GetAllUsers() ([]models.User, error) {
    return us.UserRepository.GetAll()
}

// GetUserByID retrieves a user by its ID
func (us *UserService) GetUserByID(id uint) (*models.User, error) {
    return us.UserRepository.GetByID(id)
}

// UpdateUser updates an existing user
func (us *UserService) UpdateUser(user *models.User) error {
    return us.UserRepository.Update(user)
}

// DeleteUser deletes a user
func (us *UserService) DeleteUser(user *models.User) error {
    return us.UserRepository.Delete(user)
}
