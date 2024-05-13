package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type EventRepository struct {
    db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
    return &EventRepository{db}
}

func (r *EventRepository) GetAllEvents() ([]models.Event, error) {
    var events []models.Event
    if err := r.db.Find(&events).Error; err != nil {
        return nil, err
    }
    return events, nil
}

func (r *EventRepository) GetEventByID(id uint) (*models.Event, error) {
    var event models.Event
    if err := r.db.First(&event, id).Error; err != nil {
        return nil, err
    }
    return &event, nil
}

// Implement other methods like CreateEvent, UpdateEvent, DeleteEvent as needed
