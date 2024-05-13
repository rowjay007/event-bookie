package service

import (
    "github.com/rowjay007/event-bookie/internal/models"
    "github.com/rowjay007/event-bookie/internal/repository"
)

type UserService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
    return &UserService{userRepo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.userRepo.GetAllUsers()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    return s.userRepo.GetUserByID(id)
}

// Implement other methods like CreateUser, UpdateUser, DeleteUser as needed
