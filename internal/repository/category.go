package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (cr *CategoryRepository) Create(category *models.Category) error {
	return cr.DB.Create(category).Error
}

func (cr *CategoryRepository) GetAll(queryParams map[string]string, offset, limit int) ([]models.Category, int64, error) {
	var categories []models.Category
	var total int64

	query := cr.DB.Model(&models.Category{})

	if name := queryParams["name"]; name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
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

	if offset >= 0 && limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (cr *CategoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := cr.DB.First(&category, id).Error
	return &category, err
}

func (cr *CategoryRepository) Update(category *models.Category) error {
	return cr.DB.Save(category).Error
}

func (cr *CategoryRepository) Delete(category *models.Category) error {
	return cr.DB.Delete(category).Error
}
