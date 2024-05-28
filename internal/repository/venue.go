package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type VenueRepository struct {
	DB *gorm.DB
}

func NewVenueRepository(db *gorm.DB) *VenueRepository {
	return &VenueRepository{DB: db}
}

func (vr *VenueRepository) Create(venue *models.Venue) error {
	return vr.DB.Create(venue).Error
}

func (vr *VenueRepository) GetAll(params map[string]string, offset, limit int) ([]models.Venue, int64, error) {
	var venues []models.Venue
	var total int64

	query := vr.DB.Model(&models.Venue{})

	if name, exists := params["name"]; exists {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if location, exists := params["location"]; exists {
		query = query.Where("location LIKE ?", "%"+location+"%")
	}

	if sortBy, exists := params["sort_by"]; exists {
		order := "ASC"
		if sortOrder, exists := params["sort_order"]; exists && sortOrder == "desc" {
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

	if err := query.Find(&venues).Error; err != nil {
		return nil, 0, err
	}

	return venues, total, nil
}

func (vr *VenueRepository) GetByID(id uint) (*models.Venue, error) {
	var venue models.Venue
	if err := vr.DB.First(&venue, id).Error; err != nil {
		return nil, err
	}
	return &venue, nil
}

func (vr *VenueRepository) Update(venue *models.Venue) error {
	return vr.DB.Save(venue).Error
}

func (vr *VenueRepository) Delete(id uint) error {
	return vr.DB.Delete(&models.Venue{}, id).Error
}
