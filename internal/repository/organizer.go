package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type OrganizerRepository struct {
	DB *gorm.DB
}

func NewOrganizerRepository(db *gorm.DB) *OrganizerRepository {
	return &OrganizerRepository{DB: db}
}

func (or *OrganizerRepository) Create(organizer *models.Organizer) error {
	return or.DB.Create(organizer).Error
}

func (or *OrganizerRepository) GetAll(params map[string]string, offset, limit int) ([]models.Organizer, int64, error) {
	var organizers []models.Organizer
	var total int64

	query := or.DB.Model(&models.Organizer{})

	// Filtering
	if name, exists := params["name"]; exists && name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if contact, exists := params["contact"]; exists && contact != "" {
		query = query.Where("contact LIKE ?", "%"+contact+"%")
	}

	// Sorting
	if sortBy, exists := params["sort_by"]; exists && sortBy != "" {
		order := "ASC"
		if sortOrder, exists := params["sort_order"]; exists && sortOrder == "desc" {
			order = "DESC"
		}
		query = query.Order(sortBy + " " + order)
	}

	// Count total before pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	if offset >= 0 && limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	// Execute the query
	if err := query.Find(&organizers).Error; err != nil {
		return nil, 0, err
	}

	return organizers, total, nil
}

func (or *OrganizerRepository) GetByID(id uint) (*models.Organizer, error) {
	var organizer models.Organizer
	if err := or.DB.First(&organizer, id).Error; err != nil {
		return nil, err
	}
	return &organizer, nil
}

func (or *OrganizerRepository) Update(organizer *models.Organizer) error {
	return or.DB.Save(organizer).Error
}

func (or *OrganizerRepository) Delete(id uint) error {
	return or.DB.Delete(&models.Organizer{}, id).Error
}
