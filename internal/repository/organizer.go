package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type OrganizerRepository struct {
    db *gorm.DB
}

func NewOrganizerRepository(db *gorm.DB) *OrganizerRepository {
    return &OrganizerRepository{db}
}

func (r *OrganizerRepository) GetAllOrganizers() ([]models.Organizer, error) {
    var organizers []models.Organizer
    if err := r.db.Find(&organizers).Error; err != nil {
        return nil, err
    }
    return organizers, nil
}

func (r *OrganizerRepository) GetOrganizerByID(id uint) (*models.Organizer, error) {
    var organizer models.Organizer
    if err := r.db.First(&organizer, id).Error; err != nil {
        return nil, err
    }
    return &organizer, nil
}

// Implement other methods like CreateOrganizer, UpdateOrganizer, DeleteOrganizer as needed
