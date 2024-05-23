package repository

import (
    "github.com/rowjay007/event-bookie/internal/models"
    "gorm.io/gorm"
)

// UserRepository handles user-related database operations
type UserRepository struct {
    DB *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

// Create creates a new user in the database
func (ur *UserRepository) Create(user *models.User) error {
    return ur.DB.Create(user).Error
}

// GetAll retrieves all users from the database
func (ur *UserRepository) GetAll() ([]models.User, error) {
    var users []models.User
    if err := ur.DB.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

// GetByID retrieves a user by its ID
func (ur *UserRepository) GetByID(id uint) (*models.User, error) {
    var user models.User
    if err := ur.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// Update updates an existing user in the database
func (ur *UserRepository) Update(user *models.User) error {
    return ur.DB.Save(user).Error
}

// Delete deletes a user from the database
func (ur *UserRepository) Delete(user *models.User) error {
    return ur.DB.Delete(user).Error
}
