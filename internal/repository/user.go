package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

// UserRepository handles database operations for User model
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser creates a new user record in the database
func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.DB.Create(user).Error
}

// GetUserByID retrieves a user record by its ID from the database
func (ur *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user record in the database
func (ur *UserRepository) UpdateUser(user *models.User) error {
	return ur.DB.Save(user).Error
}

// DeleteUser deletes a user record from the database
func (ur *UserRepository) DeleteUser(user *models.User) error {
	return ur.DB.Delete(user).Error
}
