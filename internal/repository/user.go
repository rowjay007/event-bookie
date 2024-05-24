package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) GetAll(queryParams map[string]string, offset, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := ur.DB.Model(&models.User{})

	if name := queryParams["name"]; name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if email := queryParams["email"]; email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	if sortBy := queryParams["sort_by"]; sortBy != "" {
		order := "ASC"
		if sortOrder := queryParams["sort_order"]; sortOrder == "desc" {
			order = "DESC"
		}
		query = query.Order(sortBy + " " + order)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if offset != -1 && limit != -1 {
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (ur *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := ur.DB.First(&user, id).Error
	return &user, err
}

func (ur *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (ur *UserRepository) Update(user *models.User) error {
	return ur.DB.Save(user).Error
}

func (ur *UserRepository) Delete(user *models.User) error {
	return ur.DB.Delete(user).Error
}
