package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type VenueRepository struct {
    db *gorm.DB
}

func NewVenueRepository(db *gorm.DB) *VenueRepository {
    return &VenueRepository{db}
}

func (r *VenueRepository) GetAllVenues() ([]models.Venue, error) {
    var venues []models.Venue
    if err := r.db.Find(&venues).Error; err != nil {
        return nil, err
    }
    return venues, nil
}

func (r *VenueRepository) GetVenueByID(id uint) (*models.Venue, error) {
    var venue models.Venue
    if err := r.db.First(&venue, id).Error; err != nil {
        return nil, err
    }
    return &venue, nil
}

// Implement other methods like CreateVenue, UpdateVenue, DeleteVenue as needed
